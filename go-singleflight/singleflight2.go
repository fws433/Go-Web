package main

import (
	"errors"
	"log"
	"sync"

	"golang.org/x/sync/singleflight"
)
//模拟10个并发请求，来同时调用getData函数，先尝试从cache中获取，如果cache中不存在就从db中获取
var g singleflight.Group
var errorNotExist=errors.New("not exist")
func main(){
	var wg sync.WaitGroup
	wg.Add(10)
	//模拟10个并发
	for i:=0;i<10;i++{
		go func(){
			defer wg.Done()
			data,err:=getData("key")  //通过getData(key)获取逻辑
			if err!=nil{
				log.Print(err)
				return
			}
			log.Println(data)
		}()
	}
	wg.Wait()
}
//获取数据
func getData(key string) (string, error) {
	data,err:=getDataFromCache(key)
	if err==errorNotExist{
		//模拟从db中获取数据
		//data,err=getDataFromDB(key)

		v,err,_:=g.Do(key,func()(interface{},error){
			return getDataFromDB(key)

		})
		if err!=nil{
			log.Println(err)
			return "",err
		}
		data=v.(string)
	}else if err!=nil{
		return "",err
	}
	return data,nil
}
//模拟从cache中获取值，cache中无该值
func getDataFromCache(key string)(string,error){
	return "",errorNotExist
}
//模拟从数据中获取值
func getDataFromDB(key string)(string,error){
	log.Printf("get %s from database",key)
	return "data",nil
}
