package main

import "net/http"

func main() {
	//当请求到达就会执行函数体内的内容
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hell World"))
	})
	//设置web服务器,第一个参数是监听地址，第二个参数是路由
	http.ListenAndServe(":8080", nil)
}
