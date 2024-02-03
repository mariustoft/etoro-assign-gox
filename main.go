package main

import (
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.New("index").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
		<title>Go Web App</title>
		</head>
		<body>
		<h1>Welcome to Go Web App</h1>
		</body>
		</html>
		`))
		tmpl.Execute(w, nil)
	})

	http.ListenAndServe(":8080", nil)
}
