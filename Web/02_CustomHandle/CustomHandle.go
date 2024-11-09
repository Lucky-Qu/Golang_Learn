package main

import "net/http"

// 创建两个自定义Handle
type helloHandle struct {
}
type aboutHandle struct {
}

// 实现Handler接口
func (helloHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello there"))
}

func (aboutHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About me"))
}

func main() {
	//新建服务器
	server := http.Server{Addr: "localhost:8080", Handler: nil}
	//注册Handle
	http.Handle("/hello", helloHandle{})
	http.Handle("/about", aboutHandle{})
	//开启服务器
	server.ListenAndServe()
}
