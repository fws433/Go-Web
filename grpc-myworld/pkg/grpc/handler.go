package grpc
//传输层，可以使用不同的通信方式，本项目中使用grpc
import (


	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/grpc"
	"grpc-world/pkg/encode"
	"grpc-world/pkg/grpc/pb"

	ep "Demo-project/grpc-myworld/pkg/endpoint"
	"golang.org/x/net/context"

)

type grpcServer struct{
	get grpc.Handler
	put grpc.Handler
}
//当使用grpc调用Get时,它只用作传递数据,当调用grpcServer中对应当handler时，ServerGRPC将其转交给go-kit处理
func(g *grpcServer)Get(ctx context.Context,req *pb.GetRequest)(*pb.ServiceResponse,error){
	_, rep, err := g.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ServiceResponse), nil
}

func (g *grpcServer) Put(ctx context.Context, req *pb.GetRequest) (*pb.ServiceResponse, error) {
	_, rep, err := g.put.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ServiceResponse), nil
}
func MakeGRPCHandler(eps ep.Endpoints, opts ...grpc.ServerOption) pb.ServiceServer {
	return &grpcServer{
		get: grpc.NewServer(
			endpoint.Endpoint(eps.GetEndpoint),
			decodeGetRequest,
			encodeResponse,
			opts...,

		),
		put: grpc.NewServer(
			endpoint.Endpoint(eps.PutEndpoint),
			decodeGetRequest,
			encodeResponse,
			opts...,
		),
	}
}

func decodeGetRequest(_ context.Context, r interface{}) (interface{}, error) {
	return ep.GetRequest{
		Key: r.(*pb.GetRequest).Key,
		Val: r.(*pb.GetRequest).Val,
	}, nil
}

func encodeResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(encode.Response)
	var errStr, data string
	var err error
	if resp.Error != nil {
		errStr = resp.Error.Error()
		err = resp.Error
	}
	if resp.Data != nil {
		data = resp.Data.(string)
	}
	return &pb.ServiceResponse{
		Success: resp.Success,
		Code:    int64(resp.Code),
		Data:    data,
		Err:     errStr,
	}, err
}




