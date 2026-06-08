package main

import (
	"html/template"
	"log"
	"net/http"
)

// tpl, err := template.Template(template.ParseFiles("templates/index.html"))

type PageData struct {
	UserText string
	Banner   string
	ASCIIArt string
}

func HomePage(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.Error(res, "Not Found", http.StatusNotFound)
		return
	}
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(res, "Error parsing a file", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(res, nil)
	if err != nil {
		http.Error(res, "Error parsing a file", http.StatusInternalServerError)
		return
	}
}

func ASCIIPage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/ascii-art" {
		http.Error(res, "Not Found", http.StatusNotFound)
		return
	}
	err := req.ParseForm()
	if err != nil {
		http.Error(res, "Error with form content", http.StatusInternalServerError)
		return
	}
	input := req.FormValue("text")
	banner := req.FormValue("banner")

	bannerMap, err := MakeBanner(banner)
	if err != nil {
		http.Error(res, "Error loading banner", http.StatusInternalServerError)
		return
	}
	Art, err := GenerateArt(input, bannerMap)
	if err != nil {
		http.Error(res, "Error Generating ASCII Art", http.StatusInternalServerError)
		return
	}
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(res, "500 Internal Server Error: Template Missing", http.StatusInternalServerError)
		return
	}
	data := PageData{
		UserText: input,
		Banner:   banner,
		ASCIIArt: Art,
	}
	tpl.Execute(res, data)
}
func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/ascii-art", ASCIIPage)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
