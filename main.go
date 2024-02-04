package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// declare variable message
		// message := "Hello World"
		phone := r.Form.Get("hone")
		log.Println("phone-->", phone)

		root, _ := template.ParseFiles("index.html")
		// wallet, _ := template.ParseFiles("wallet.html")
		// add variable to template
		root.Execute(w, "message")

	})

	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		// declare variable message
		message := "Hello World"
		// add variable to template
		root, _ := template.ParseFiles("index.html")
		root.Execute(w, message)

	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
