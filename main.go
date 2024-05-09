package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy struct { //定义了一个名为Proxy的空结构体，它将用作http.Handler的实现
}

// ServeHTTP方法是http.Handler接口的一部分,实现ListenAndServe()第二个参数里适配器的接口，作为适配器Proxy{}传入
func (Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, _ := url.Parse("http://127.0.0.1:8082")
	reverseProxy := httputil.NewSingleHostReverseProxy(remote)
	reverseProxy.ServeHTTP(w, r)
}

func main() {
	addr := "127.0.0.1:8081"

	fmt.Printf("反向代理正在监听本地端口：%s\n", addr)
	err := http.ListenAndServe(addr, Proxy{})
	if err != nil {
		fmt.Println(err)
	}
}
