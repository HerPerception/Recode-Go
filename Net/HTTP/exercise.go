package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(write http.ResponseWriter, read *http.Request) {
	if read.URL.Path != "/" {
		http.Error(write, "404 Not Found", http.StatusNotFound)
		return
	}
	if read.Method != "GET" {
		http.Error(write, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	htmlContent := `
	<!DOCTYPE html>
	<html>
	<head>
	<title>First Exercise</title>
	</head>
	<body>
	<h3 style="background-color: pink">"Hello there! This is the first exercise on the web part."</h3>
	<a href="/submit" target=_self style="background-color: yellow; color: black;">submit</a>
	</body>
	</html>
	`
	fmt.Fprint(write, htmlContent)
}

func SubmitHandler(write http.ResponseWriter, read *http.Request) {
	if read.URL.Path != "/submit" {
		http.Error(write, "404 Not Found", http.StatusNotFound)
		return
	}
	// if read.Method != "POST" {
	// 	http.Error(write, "405 Method Not Allowed", http.StatusMethodNotAllowed)
	// 	return
	//  }
	fmt.Fprint(write, "This is the submit path.")
}
func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/submit", SubmitHandler)
	fmt.Println("Started server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
