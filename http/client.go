package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	// resp, err := http.DefaultClient.Do(request)//默认Client
	//默认方法发请求
	// resp, err := http.Get("http://www.imooc.com") //发送Get请求
	//（option）自定义请求：指定Header发送请求
	request, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Mobile Safari/537.36") //设置移动端的User-Agent
	// 自己构造client
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { //判断是否发生了重定向，重定向的路径放在via []*http.Request
			fmt.Println("Redirect: ", req)
			// fmt.Println("via: ", via)
			return nil
		}, //如果CheckRedirect非空，则HTTP重定向前client会调用CheckRedirect
	}
	resp, err := client.Do((request))

	if err != nil { //错误处理
		panic(err)
	}
	defer resp.Body.Close() //需要关闭连接

	//解析响应resp
	s, err := httputil.DumpResponse(resp, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
