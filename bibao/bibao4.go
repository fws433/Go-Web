package main

import "fmt"
//难以理解，，=有可能是头昏，下次在看
func main(){
	var flist []func()
	for i:=0;i<3;i++{
		flist=append(flist, func() {
			fmt.Println(i)
		})
	}
	for _,f:=range flist{
		f()
	}
}
