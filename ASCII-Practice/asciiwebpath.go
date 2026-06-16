package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Text string
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error parsing file", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error parsing file", http.StatusInternalServerError)
		return
	}
	//fmt.Fprintln(w, "Hello from my server")
}
func HomeNew(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/echo" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	// tpl, err := template.ParseFiles("index.html")
	// if err != nil {
	// 	http.Error(w, "Error parsing file", http.StatusInternalServerError)
	// 	return
	// }
	text := r.FormValue("text")
	// Data := Page{
	// 	Text: text,
	// }
	// err = tpl.Execute(w, Data)
	// if err != nil {
	// 	http.Error(w, "Error parsing file", http.StatusInternalServerError)
	// 	return
	// }
	fmt.Fprintln(w, text)
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/echo", HomeNew)
	http.ListenAndServe(":8080", nil)
}
