package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HomePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.Error(res, "404 Not found", http.StatusNotFound)
		return
	}

	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(res, "500 Error parsing files", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(res, nil)
	if err != nil {
		http.Error(res, "500 Error with files", http.StatusInternalServerError)
		return
	}
}

func ASCIIHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/ascii-art" {
		http.Error(res, "Not found", http.StatusNotFound)
		return
	}
	if req.Method != "POST" {
		http.Error(res, "400 method not allowed", http.StatusBadRequest)
		return
	}
	bans := req.FormValue("banner")
	text := req.FormValue("text")
	fmt.Println(text)
	fmt.Println(bans)
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(res, "500 Error parsing files", http.StatusInternalServerError)
		return
	}
	banner := fmt.Sprintf("banners/%s", bans)
	art, err := GenerateArt(text, banner)

	if err != nil {
		http.Error(res, "500 Error generating ASCII Art.", http.StatusInternalServerError)
		return
	}
	finalResult := map[string]string{
		"UserText": text,
		"Result":   art,
	}
	err = tpl.ExecuteTemplate(res, "index.html", finalResult)
	if err != nil {
		http.Error(res, "500 Error generating ASCII Art.", http.StatusInternalServerError)
		return
	}
}
func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/ascii-art", ASCIIHandler)
	fmt.Println("Server starts at port :8080 on localhost")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
