package main

import (
	book "Demo-project/go-book-grpc"
	"context"
	endpoint2 "github.com/go-kit/kit/endpoint"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"kit/endpoint"
	"log"
	"net"
)
//type BookServer struct{}
type BookServer struct{
	bookListHandler grpc_transport.Handler
	bookInfoHandler grpc_transport.Handler

}

//通过grpc调用GetBookInfo时，GetBookInfo只用作传递数据，当调用bookserver中对应的handler.servergrpc转交给go-kit处理
func(s *BookServer)GetBookInfo(ctx context.Context,in *book.BookInfoParams)(*book.BookInfo,error){
	_,rsp,err:=s.bookInfoHandler.ServeGRPC(ctx,in)
	if err!=nil{
		return nil,err
	}
	return rsp.(*book.BookInfo),err
}

/*func(s *BookServer)GetBookInfo(ctx context.Context,in *book.BookInfoParams)(*book.BookInfo,error){
	//请求详情时返回书籍信息
	b:=new(book.BookInfo)
	b.BookId=in.BookId
	b.BookName="Go语言核心编程"
	return b,nil
}*/

func(s *BookServer)GetBookList(ctx context.Context,in *book.BookListParams)(*book.BookList,error){
	_,rsp,err:=s.bookInfoHandler.ServeGRPC(ctx,in)
	if err!=nil{
		return nil,err
	}
	return rsp.(*book.BookList),err
}

/*func(s *BookServer)GetBookList(ctx context.Context,in *book.BookListParams)(*book.BookList,error){
	//请求列表时返回书籍列表
	bl:=new(book.BookList)
	bl.BookList=append(bl.BookList,&book.BookInfo{BookId: 1,BookName: "Go语言核心编程"})
	bl.BookList=append(bl.BookList,&book.BookInfo{BookId: 2,BookName: "java从入门到精通"})
	return bl,nil
}*/

//创建bookInfo的endpoint
//定义好函数后，把函数注册到http或者grpc上，就可以实现业务函数到回调
func makeGetBookInfoEndpoint() endpoint.Endpoint{
	return func(ctx context.Context,request interface{})(interface{},error){
		//请求详情时返回书籍信息
		req:=request.(*book.BookInfoParams)
		b:=new(book.BookInfo)
		b.BookId=req.BookId
		b.BookName="Go-kit与微服务"
		return b ,nil
	}
}
//创建bookList的EndPoint
func makeGetBookListEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//请求列表时返回 书籍列表
		bl := new(book.BookList)
		bl.BookList = append(bl.BookList, &book.BookInfo{BookId: 1, BookName: "Go入门到精通"})
		bl.BookList = append(bl.BookList, &book.BookInfo{BookId: 2, BookName: "go-kit入门到精通"})
		bl.BookList = append(bl.BookList, &book.BookInfo{BookId: 2, BookName: "java入门到精通"})
		return bl, nil
	}
}
func decodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}


func main(){
	//包装BookServer
	bookServer:=new(BookServer)
	//创建booklist的handler
	bookListHandler:=grpc_transport.NewServer(
		endpoint2.Endpoint(makeGetBookListEndpoint()),
		decodeRequest,
		encodeResponse,
		)
	//bookserver 增加go-kit流程的bookLIST处理逻辑
	bookServer.bookListHandler=bookListHandler
	//创建bookInfo的Handler
	bookInfoHandler := grpc_transport.NewServer(
		endpoint2.Endpoint(makeGetBookInfoEndpoint()),
		decodeRequest,
		encodeResponse,
	)
	//bookServer 增加 go-kit流程的 bookInfo处理逻辑
	bookServer.bookInfoHandler = bookInfoHandler
	//启动grpc服务
	serviceAddress:=":50053"

	//创建tcp监听
	lis,err:=net.Listen("tcp",serviceAddress)
	if err!=nil{
		log.Fatalf("failed to listen:%v",err)

	}
	//创建grpc服务
	grpcServer:=grpc.NewServer()
	//注册bookServer
	book.RegisterBookServiceServer(grpcServer,bookServer)
	//启动服务
	grpcServer.Serve(lis)
}


