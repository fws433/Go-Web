package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

//爬虫步骤：
/*
明确目标
爬(爬下内容)
取(筛选想要的)
处理数据（按照你的的想法去处理）


*/
import (
	"fmt"
)
var(
	reQQEmail=`(\d+)@qq.com`
)

//爬邮箱
func main(){
	GetEmail()
}

func GetEmail() {
	//1去网站拿数据
	resp,err:=http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731" )
	HandleError(err,"http.Get url")
	defer resp.Body.Close()
	//2.读取页面内容
	pageBytes,err:=ioutil.ReadAll(resp.Body)
	HandleError(err,"ioutil.Readall")
	//字节转字符串
	pageStr:=string(pageBytes)
	//3.过滤数据，过滤qq邮箱
	re:=regexp.MustCompile(reQQEmail)
	//-1代表全部
	results:=re.FindAllStringSubmatch(pageStr,-1)
	//遍历结果
	for _,result:=range results{
		fmt.Println("email:",result[0])
		fmt.Println("qq:",result[1])
	}

}
//处理异常
func HandleError(err interface{}, s string) {
	if err!=nil{
		fmt.Println(s,err)
	}
}
