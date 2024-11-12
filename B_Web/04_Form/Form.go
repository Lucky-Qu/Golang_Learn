package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("B_Web/04_Form/File")))
	http.HandleFunc("/File", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			fmt.Fprintln(w, r.Form)
		}
	})
	http.ListenAndServe(":8080", nil)
}
