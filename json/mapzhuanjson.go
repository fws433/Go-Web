package main

import (
	"encoding/json"
	"fmt"
)

/*type User1 struct{
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age     int      `json:"age"`
	Birthday string `json:"birthday"`
	Sex string      `json:"sex"`
}*/
func testMap(){
	var mmp map[string]interface{}
	mmp=make(map[string]interface{})
	mmp["username"]="murpjy"
	mmp["age"]=19
	mmp["sex"]="man"
	data,err:=json.Marshal(mmp)
	if err!=nil{
		fmt.Println("json marshal failed,err:",err)
		return
	}
	fmt.Printf("%s\n",string(data))

}
func main(){
	testMap()
	fmt.Println("....")
}
