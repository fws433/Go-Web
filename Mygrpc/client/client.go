package main

import (
	pb "Demo-project/Mygrpc"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main(){
	//客户端连接服务器
	conn,err:=grpc.Dial("127.0.0.1:50054",grpc.WithInsecure())
	if err!=nil{
		fmt.Println("网络异常",err)
	}
	defer conn.Close()
	//获得grpc句柄
	c:=pb.NewHelloserverClient(conn)
	//通过句柄调用函数
	re,err:=c.Sayhello(context.Background(),&pb.HelloReq{Name:"熊猫"})
	if err!=nil{
		fmt.Println("sayHello服务调用失败")
	}
	fmt.Println("调用sayHello的返回：",re.Msg)

	rel,err:=c.Sayname(context.Background(),&pb.NameRep{Name: "托马斯"})
	if err!=nil{
		fmt.Println("sayName服务调用失败")
	}
	fmt.Println("调用sayName的返回：",rel.Name)

}
