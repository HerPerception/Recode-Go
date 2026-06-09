// This file is to learn exclusively the go net/http and html/template
package main

import (
	"html/template"
	"net/http"
)

type MyDatabase struct {
	URL string
}

// Implement the Handler interface
func (db *MyDatabase) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Connecting to: " + db.URL))
}

// func main() {
// 	dbHandler := &MyDatabase{URL: "postgres://..."}

// 	// Pass the object directly into http.Handle
// 	http.Handle("/db-status", dbHandler)
// 	http.ListenAndServe(":8080", nil)
// }

func ParseHome(write http.ResponseWriter, read *http.Request) {
	tpl, err := template.ParseFiles("templates/idex2.html", "templates/index.html")
	if err != nil {
		http.Error(write, "Error parsing files", http.StatusInternalServerError)
		return
	}
	err = tpl.ExecuteTemplate(write, "index.html", nil)
	if err != nil {
		http.Error(write, "Error sending a file", http.StatusInternalServerError)
		return
	}
}

// func main() {
// 	http.HandleFunc("/", ParseHome)
// 	err := http.ListenAndServe(":8080", nil)
// 	log.Fatal(err)
// }
