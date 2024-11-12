package main

import (
	"html/template"
	"net/http"
)

func testTemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("B_Web/05_Template/index.html")
	t.Execute(w, "Hello There")
}

func main() {
	server := http.Server{Addr: "localhost:8080", Handler: nil}
	http.HandleFunc("/test", testTemplate)
	server.ListenAndServe()
}
