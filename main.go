package main

import (
	"html/template"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// templ, _ := template.New("test").Parse("Hello, {{.}}!")
		// add html file
		templ, _ := template.ParseFiles("index.html")

		templ.Execute(w, "")

	})
	
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
