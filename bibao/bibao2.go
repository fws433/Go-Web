package main

import "fmt"
//闭包=引用环境+匿名函数
//函数square每进入一次，就形成了一个新的环境
//闭包是地址引用
func squares() func()int{
	var x int
	return func() int{  //匿名函数
		x++
		return x*x
	}
}
func main(){
	f1:=squares()// 引用环境
	f2:=squares()//引用环境
	fmt.Println("first call f1:",f1())
	fmt.Println("second call f1:",f1())
	fmt.Println("first call f2:",f2())
	fmt.Println("second call f2:",f2())

}
