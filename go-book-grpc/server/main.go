package main

import (
	book "Demo-project/go-book-grpc"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)
/*方法是导出的
方法有两个参数，都是导出类型或内建类型
方法的第二个参数是指针
 方法只有一个error接口类型的返回值*/

//服务器用于实现BookService
type BookServer struct{}
//函数关键字（对象）函数名（默认是传参，客户端传入的参数）（服务端返回的参数，错误返回值）
func(s *BookServer)GetBookInfo(ctx context.Context,in *book.BookInfoParams)(*book.BookInfo,error){
	//请求详情时返回书籍信息
	b:=new(book.BookInfo)
	b.BookId=in.BookId
	b.BookName="Go语言核心编程"
	return b,nil
}
func(s *BookServer)GetBookList(ctx context.Context,in *book.BookListParams)(*book.BookList,error){
	//请求列表时返回书籍列表
	bl:=new(book.BookList)
	bl.BookList=append(bl.BookList,&book.BookInfo{BookId: 1,BookName: "Go语言核心编程"})
	bl.BookList=append(bl.BookList,&book.BookInfo{BookId: 2,BookName: "java从入门到精通"})
	return bl,nil
}

func main(){
	//启动grpc服务
	serviceAddress:=":50053"
	//将类实力话为对象
	bookServer:=new(BookServer)
	//创建tcp监听
	lis,err:=net.Listen("tcp",serviceAddress)
	if err!=nil{
		log.Fatalf("failed to listen:%v",err)

	}
	//创建grpc服务
	grpcServer:=grpc.NewServer()
	//服务端注册个对象，注册bookServer
	book.RegisterBookServiceServer(grpcServer,bookServer)
	//启动服务
	grpcServer.Serve(lis)
}
