package service

import (
	ep "Demo-project/grpc-myworld/pkg/endpoint"
	"Demo-project/grpc-myworld/pkg/grpc"
	"Demo-project/grpc-myworld/pkg/grpc/pb"
	"Demo-project/grpc-myworld/pkg/repository"
	"Demo-project/grpc-myworld/pkg/service"
	"flag"
	"fmt"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/oklog/pkg/group"
	"golang.org/x/time/rate"
	googleGrpc "google.golang.org/grpc"
	"kit/endpoint"
	"kit/log"
	"kit/log/level"
	"kit/transport"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//在Run里面需要初始化数据仓库Repository，
//初发化Service，初始化Endpoint和初始化Transport，初始化完成之后再启动相用grpc个传输方式。
var logger log.Logger

const rateBucketNum = 20

var (
	fs       = flag.NewFlagSet("world", flag.ExitOnError)
	httpAddr = fs.String("http-addr", ":8080", "HTTP listen address")
	grpcAddr = fs.String("grpc-addr", ":8081", "gRPC listen address")
)

func Run() {

	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	store := repository.New()

	svc := service.New(logger, store)
	svc = service.NewLoggingService(logger, svc)

	ems := []endpoint.Middleware{
		service.TokenBucketLimitter(rate.NewLimiter(rate.Every(time.Second*1), rateBucketNum)), // 限流
	}

	eps := ep.NewEndpoint(svc, map[string][]endpoint.Middleware{
		"Get": ems,
		"Put": ems,
	})

	g := &group.Group{}
	initGRPCHandler(eps, g)
	initCancelInterrupt(g)

	_ = level.Error(logger).Log("exit", g.Run())
}

func initCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}


func initGRPCHandler(endpoints ep.Endpoints, g *group.Group) {
	grpcOpts := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		_ = logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}

	g.Add(func() error {
		baseServer := googleGrpc.NewServer()
		pb.RegisterServiceServer(baseServer, grpc.MakeGRPCHandler(endpoints, grpcOpts...))
		return baseServer.Serve(grpcListener)
	}, func(error) {
		_ = level.Error(logger).Log("grpcListener.Close", grpcListener.Close())
	})
}