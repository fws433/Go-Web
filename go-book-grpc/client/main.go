package main

import (
	book "Demo-project/go-book-grpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main(){
	serviceAddress:="127.0.0.1:50053"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()
	//调用protobuf的函数创建客户端具柄
	bookClient := book.NewBookServiceClient(conn)
	//调用protobuf写好的函数
	bi, _ := bookClient.GetBookInfo(context.Background(), &book.BookInfoParams{BookId: 1})
	fmt.Print("获取书籍信息:\t")
	fmt.Println("bookInfo:", bi.BookName)

	bl, _ := bookClient.GetBookList(context.Background(), &book.BookListParams{Page: 1, Limit: 10})
	fmt.Println("获取书籍列表:\t")
	for _, b := range bl.BookList {
		fmt.Println("bookId:", b.BookId, " => ", "bookName:", b.BookName)
	}


}
