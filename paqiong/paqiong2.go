package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)
//需求：给一个网站，爬取邮箱，超链接，手机号，身份证号，以及图片链接
//爬虫步骤：
/*
1.明确目标（确定在哪个url上爬去）
2.爬下网页内容，并将字节转成字符串
3.筛选（根据全局变量中定义的正则表达式规则进行过滤）
4.处理数据
*/
var(
	//w代表大小写字母+数字+下划线
	reEmail=`\w+@\w+\.\w+`
	reLinke  = `href="(https?://[\s\S]+?)"`
	rePhone  = `1[3456789]\d\s?\d{4}\s?\d{4}`
	reIdcard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
	reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)
func main(){
	// 2.抽取的爬邮箱
	// GetEmail2("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	// 3.爬链接
	//GetLink("http://www.hao123.com")
	// 4.爬手机号
	//GetPhone("https://www.zhaohaowang.com/")
	// 5.爬身份证号
	//GetIdCard("https://henan.qq.com/a/20171107/069413.htm")
	// 6.爬图片
	GetImg("http://image.baidu.com/search/index?tn=baiduimage&ps=1&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=%E7%BE%8E%E5%A5%B3")
}
//3爬图片
func GetImg(url string) {
	pageStr:=GetPageStr(url)
	re:=regexp.MustCompile(reImg)
	results:=re.FindAllStringSubmatch(pageStr,-1)
	for _,result:=range results{
		fmt.Println(result[0])
	}
}

//3爬身份证号
/*func GetIdCard(url string) {
	pageStr:=GetPageStr(url)
	re:=regexp.MustCompile(reIdcard)
	results:=re.FindAllStringSubmatch(pageStr,-1)
	for _,result:=range results{
		fmt.Println(result[0])
	}
}*/

//3.爬手机号
/*func GetPhone(url string) {
	pageStr:=GetPageStr(url)
	//过滤数据，过滤身份证号
	re:=regexp.MustCompile(rePhone)
	results:=re.FindAllStringSubmatch(pageStr,-1)
	for _,result:=range results{
		fmt.Println(result[0])
	}
}*/

//3.爬链接
/*func GetLink(url string) {
	pageStr:=GetPageStr(url)
	re:=regexp.MustCompile(reLinke)
	results:=re.FindAllStringSubmatch(pageStr,-1)
	for _,result:=range  results{
		fmt.Println(result[1])
	}

}*/

//根据url获取内容
func GetPageStr(url string) (pageStr string) {
	resp,err:=http.Get(url)
	HandleError1(err,"http.Get url")
	defer resp.Body.Close()
	//2.读取页面内容
	pageBytes,err:=ioutil.ReadAll(resp.Body)
	HandleError1(err,"ioutil.readall")
	//字节转字符串
	pageStr=string(pageBytes)
	return pageStr

}

func HandleError1(err error, s string) {
	if err!=nil{
		fmt.Println(s,err)
	}
}