package main

import "net/http"

func main() {
	//使用FileServer委托Dir使用FileTest内的文件处理请求
	http.ListenAndServe("localhost:8080", http.FileServer(http.Dir("Web/03_FileServer/FileTest")))
}
