package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 浏览器其实就是一个发送和接收HTTP协议数据的客户端，
// 我们平时通过浏览器访问网页其实就是从网站的服务器接收HTTP数据，然后浏览器会按照HTML、CSS等规则将网页渲染展示出来
func main() {

	apiUrl := "https://www.liwenzhou.com//get"
	// URL param
	data := url.Values{}
	data.Set("name", "小王子")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode() // URL encode
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	// fmt.Println(string(b))

	// resp, err := http.Get("https://www.liwenzhou.com/")
	// if err != nil {
	// 	fmt.Printf("get failed,err:%v\n", err)
	// 	return
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("read from resp.body failed,err:%v\n", err)
	// 	return
	// }
	// fmt.Println(string(body))
}
