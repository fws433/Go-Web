package main
import (
	pb "Demo-project/Mygrpc"
	"context"
	endpoint2 "github.com/go-kit/kit/endpoint"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"kit/endpoint"
	"log"
	"net"
)
/*方法是导出的
方法有两个参数，都是导出类型或内建类型
方法的第二个参数是指针
 方法只有一个error接口类型的返回值*/

//服务器用于实现BookService
//type Server struct{}
type Server struct{
	sayhellohander grpc_transport.Handler
	saynamehander grpc_transport.Handler
}
//go-kit中操作方法
//当使用grpc调用Sayhello时,它只用作传递数据，当调用Server中对应当handler时，servergrpc将其转交给go-kit处理
func(s *Server)Sayhello(ctx context.Context,in *pb.HelloReq)(*pb.HelloRsp,error){
	_,rsp,err:=s.sayhellohander.ServeGRPC(ctx,in)
	if err!=nil{
		return nil,err
	}
	return rsp.(*pb.HelloRsp),err

}
//grpc中当操作方法

//函数关键字（对象）函数名（默认是传参，客户端传入的参数）（服务端返回的参数，错误返回值）
/*func(s *Server) Sayhello(ctx context.Context, in *pb.HelloReq) ( *pb.HelloRsp,  error){
	out:=new(pb.HelloRsp)
	out.Msg="hello,fws"+"=>"+in.Name
	return out,nil

	//return &pb.HelloRsp{Msg:"hello,fws"+"=>"+in.Name},nil

}*/

//grpc中当操作方法

/*func(s *Server)Sayname(ctx context.Context, in *pb.NameRep) (out *pb.NameRsp, err error){

	return &pb.NameRsp{Name: "早上好"+"=>"+in.Name},nil

}*/
//go-kit中操作方法
func(s *Server)Sayname(ctx context.Context, in *pb.NameRep) (*pb.NameRsp, error){
	_,rsp,err:=s.saynamehander.ServeGRPC(ctx,in)
	if err!=nil{
		return nil,err
	}
	return rsp.(*pb.NameRsp),err

}

//创建Sayhello的endpoint
//定义好函数后，把函数注册到grpc上，就可以实现业务函数到回调
func makeSayhelloEndpoint() endpoint.Endpoint{
	return func(ctx context.Context,request interface{})(interface{},error){
		req:=request.(*pb.HelloReq)
		b:=new(pb.HelloRsp)
		b.Msg="hello,fws"+"=>"+req.Name
		return b,nil
	}
}
//创建SayName的Endpoint,实现各种接口的handler
func makeSaynameEndpoint() endpoint.Endpoint{
	return func(ctx context.Context,request interface{})(interface{},error){
		req:=request.(*pb.NameRep)
		b:=new(pb.NameRsp)
		b.Name="早上好"+"=>"+req.Name
		return b,nil
	}
}
func decodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}
func main(){
	//启动grpc服务
	serviceAddress:=":50054"
	//将类实力话为对象
	pbServer:=new(Server)
	//创建sayhello的handler
	sayhellohander:=grpc_transport.NewServer(
		endpoint2.Endpoint(makeSayhelloEndpoint()),
		decodeRequest,
		encodeResponse,
		)
	//pbserver增加go-kit流程的sayhello逻辑
	pbServer.sayhellohander=sayhellohander

	//创建sayname的handler
	saynamehander:=grpc_transport.NewServer(
		endpoint2.Endpoint(makeSaynameEndpoint()),
		decodeRequest,
		encodeResponse,
	)
	//pbserver增加go-kit流程的saynamel逻辑
	pbServer.saynamehander=saynamehander
	//创建tcp监听
	lis,err:=net.Listen("tcp",serviceAddress)
	if err!=nil{
		log.Fatalf("failed to listen:%v",err)

	}
	//创建grpc服务
	grpcServer:=grpc.NewServer()
	//服务端注册个对象，注册bookServer
	pb.RegisterHelloserverServer(grpcServer,pbServer)
	//启动服务,等待网络连接
	grpcServer.Serve(lis)
}

