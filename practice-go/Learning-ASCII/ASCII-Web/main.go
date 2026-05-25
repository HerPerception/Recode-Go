package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Input    string
	ASCIIArt string
}

func Handler(write http.ResponseWriter, read *http.Request) {
	if read.URL.Path != "/" {
		http.Error(write, "404: Not Found", http.StatusNotFound)
		return
	}
	//fmt.Fprint(write, "This is your first class on ASCII Art Web!")
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(write, "500: Internal Server Error", http.StatusInternalServerError)
		return
	}
	tpl.Execute(write, nil)
}

func ASCIIHandler(write http.ResponseWriter, read *http.Request) {
	if read.Method != "POST" {
		http.Error(write, "405: Only POST Method Allowed", http.StatusMethodNotAllowed)
		return
	}
	err := read.ParseForm()
	if err != nil {
		http.Error(write, "400 Bad Request", http.StatusBadRequest)
		return
	}
	userInput := read.FormValue("text")
	banner := read.FormValue("banner")

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(write, "400 Bad Request:Invalid Banner", http.StatusBadRequest)
		return
	}

	bannerMap, err := LoadBanner(banner)
	if err != nil {
		http.Error(write, "500 Internal Server Error: Cannot load banner file", http.StatusInternalServerError)
		return
	}
	finalArt, err := GenerateArt(userInput, bannerMap)
	if err != nil {
		http.Error(write, "500 Internal Server Error: Problem parsing input or banner file", http.StatusInternalServerError)
		return
	}

	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(write, "500 Internal Server Error: Template Missing", http.StatusInternalServerError)
		return
	}
	data := PageData{
		Input:    userInput,
		ASCIIArt: finalArt,
	}

	tpl.Execute(write, data)

}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/ascii-art", ASCIIHandler)
	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
