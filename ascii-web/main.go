package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HomePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.Error(res, "Not found", http.StatusNotFound)
		return
	}

	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(res, "Error parsing files", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(res, nil)
	if err != nil {
		http.Error(res, "Error with files", http.StatusInternalServerError)
		return
	}
}

func ASCIIHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/ascii-art" {
		http.Error(res, "Not found", http.StatusNotFound)
		return
	}

	banner := req.FormValue("banner")
	text := req.FormValue("text")
	fmt.Println(text)
	fmt.Println(banner)
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(res, "Error parsing files", http.StatusInternalServerError)
		return
	}
	art, err := Generator(text, banner)
	// fmt.Println(req)
}
func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/ascii-art", ASCIIHandler)
	fmt.Println("Server starts at port :8080 on localhost")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
