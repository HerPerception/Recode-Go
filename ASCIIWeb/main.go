package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Global template variable to avoid parsing files on every single request
var tpl *template.Template

func init() {
	// Pre-parsing the template once when the server boots up
	var err error
	tpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("Critical Error: Failed to parse template files: %v", err)
	}
}

func HomePage(write http.ResponseWriter, read *http.Request) {
	// Fixing the catch-all issue: return 404 if the path isn't exactly "/"
	if read.URL.Path != "/" {
		http.Error(write, "404 Page Not Found", http.StatusNotFound)
		return
	}

	if read.Method != "GET" {
		http.Error(write, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	// Rendering the template with nil (empty fields) on the initial load
	err := tpl.Execute(write, nil)
	if err != nil {
		http.Error(write, "An Error occurred rendering the page", http.StatusInternalServerError)
		return
	}
}

func ASCIIPage(write http.ResponseWriter, read *http.Request) {
	if read.Method != "POST" {
		http.Error(write, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieving values submitted from the form
	textValue := read.FormValue("text")
	bannerValue := read.FormValue("banner")
	fmt.Println(textValue, bannerValue)
	// Validating that the user didn't tamper with the form values
	if bannerValue != "standard" && bannerValue != "shadow" && bannerValue != "thinkertoy" {
		http.Error(write, "400 Bad Request: Invalid banner style chosen", http.StatusBadRequest)
		return
	}

	// Executing the ASCII art generator logic
	artResult, err := Generator(textValue, bannerValue)
	if err != nil {
		// If the generator can't find a banner file, it's usually an internal issue
		http.Error(write, "Error Generating ASCII Art", http.StatusInternalServerError)
		return
	}

	// Map keys must match the case-sensitive template tags: {{.Text}} and {{.Result}}
	templateData := map[string]string{
		"Text":   textValue,
		"Result": artResult,
	}

	err = tpl.ExecuteTemplate(write, "index.html", templateData)
	if err != nil {
		http.Error(write, "An Error occurred rendering the art result", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/ascii-art", ASCIIPage)

	log.Println("Server starting on http://localhost:8080 ...")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
