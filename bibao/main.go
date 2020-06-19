package main

import "fmt"

func test() int{
	fmt.Println(2)
	return 3
}
func main(){
	c:=test()
	fmt.Println(c)
}
