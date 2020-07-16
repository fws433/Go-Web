package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

//如何并发爬取美图
var(
	//存放图片链接的数据管道
	chanImageUrls chan string
	waitGroup sync.WaitGroup
	//用于监控携程
	chanTask chan string
	reImg1=`https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`

)
func main(){
	//1.初始化管道
	chanImageUrls=make(chan string,1000000000)
	chanTask=make(chan string,26)
	//2.爬虫协程
	for i:=1;i<27;i++{
		waitGroup.Add(1)
		go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}
	//3.任务统计协程，统计26个任务是否都完成，完成则关闭管道
	waitGroup.Add(1)

	go CheckOK()

	waitGroup.Wait()

}

//任务统计协程
func CheckOK() {
	var count int
	for{
		url:=<-chanTask
		fmt.Printf("%s 完成里爬取任务\n",url)
		count++
		if count==26{
			close(chanImageUrls)
			break
		}
	}
	waitGroup.Done()
}

//爬取链接到管道
//url是传递的整页链接
func getImgUrls(url string) {
	urls:=getImg(url)
	//遍历切片里所有链接，存入数据管道
	for _,url:=range urls{
		chanImageUrls<-url
	}
	//表识当前协程完成，每完成一个任务，写一条数据，用于监控协程知道已经完成里几个任务
	chanTask<-url
	waitGroup.Done()
}

//获取当前图片链接
func getImg(url string) (urls []string) {
	pageStr:=GetPageStr1(url)
	re:=regexp.MustCompile(reImg1)
	results:=re.FindAllStringSubmatch(pageStr,-1)
	fmt.Printf("共找到%d条结果\n",len(results))
	for _,result:=range results{
		url:=result[0]
		fmt.Println(url)
		urls=append(urls,url)
	}
	//fmt.Println(urls)
	return
}

//抽取根据url获取内容
func GetPageStr1(url string) (pageStr string) {
	resp,err:=http.Get(url)
	HandleError2(err,"http.get url")
	defer resp.Body.Close()
	//读取页面信息
	pageBytes,err:=ioutil.ReadAll(resp.Body)
	HandleError2(err,"ioutil.readall")
	//字节转字符串
	pageStr=string(pageBytes)
	return pageStr
}

func HandleError2(err error, s string) {
	if err!=nil{
		fmt.Println(s,err)
	}
}

