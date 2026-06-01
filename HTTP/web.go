package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.Error(res, "404 Page Not Found", http.StatusNotFound)
		return
	}

	//	fmt.Fprintln(res, "Hello There!")
	// err := template.ParseFiles("templates/index.html")
	// if err != nil {
	// 	http.Error(res, "Internal server error", http.StatusInternalServerError)
	// 	return
	// }
	err := tpl.Execute(res, nil)
	if err != nil {
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func ASCIIHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/ascii-art" {
		http.Error(res, "Bad request", http.StatusBadRequest)
		return
	}
	err := req.ParseForm()
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}
	text := req.FormValue("text")
	bans := req.FormValue("banner")
	banner := fmt.Sprintf("%s.txt", bans)

	bannerMap, err := MakeBanner(banner)
	if err != nil {
		http.Error(res, "Error with form value", http.StatusInternalServerError)
		return
	}
	Art, err := GenerateArt(text, bannerMap)
	if err != nil {
		http.Error(res, "Might be an error generating ASCII Art", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(res, Art)
	if err != nil {
		http.Error(res, "Error parsing output", http.StatusInternalServerError)
		return
	}
}
func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii-art", ASCIIHandler)
	http.ListenAndServe(":8080", nil)
}
