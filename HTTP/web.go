package main

import (
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

type PageData struct {
	UserText string
	ASCIIArt string
}

func HomePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.Error(res, "Not Found", http.StatusNotFound)
		return
	}
	err := tpl.Execute(res, nil)
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
	in := req.FormValue("text")
	banner := req.FormValue("banner")

	bannerMap, err := MakeBanner(banner)
	if err != nil {
		http.Error(res, "Error loading banner", http.StatusInternalServerError)
		return
	}
	Art, err := GenerateArt(in, bannerMap)
	if err != nil {
		http.Error(res, "Error Generating ASCII Art", http.StatusInternalServerError)
		return
	}
	data := PageData{
		UserText: in,
		ASCIIArt: Art,
	}
	tpl.Execute(res, data)
}
func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/ascii-art", ASCIIPage)
	http.ListenAndServe(":8080", nil)
}
