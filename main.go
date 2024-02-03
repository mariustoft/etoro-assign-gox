package main

import (
	"html/template"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		templ, _ := template.New("test").Parse("Hello, {{.}}!")
		templ.Execute(w, "World")

	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
