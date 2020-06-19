package main

import "fmt"

//go语言支持闭包，go语言中如何实现闭包？
func a() func() int{
	i:=0
	b:=func() int{
		i++
		fmt.Println(i)
		return i
	}
	return b
}
func main(){
	c:=a()
	c()
	c()
	c()
	a()

}
