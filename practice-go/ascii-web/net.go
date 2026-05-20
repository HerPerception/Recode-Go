package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func helloHandler(write http.ResponseWriter, read *http.Request) {
	if read.URL.Path != "/" {
		http.Error(write, "404 Page Not Found", http.StatusNotFound)
		return
	}
	// write.Header().Set("Content-Type", "text/html; charset=utf-8")
	// fmt.Fprint(write, "<h1>Hello Medicine</h1><h2>The Only Site Where You get the Most Recent Updates on Health Matters</h2><p>How may we help you today?</p>")
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(write, "500 Internal Server Error: Internal Server Error", http.StatusInternalServerError)
		return
	}
	tpl.Execute(write, nil)
}
func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
