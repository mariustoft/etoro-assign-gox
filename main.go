package main

import (
	"html/template"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/home.html", "templates/header.html", "templates/convertor.html")

		tmpl.Execute(w, map[string]interface{}{
			"Header": r.FormValue("new-header"),
			"Portfolio": []map[string]interface{}{
				{"Name": "Cardano", "Quantity": 12},
				{"Name": "Bitcoin", "Quantity": 22},
				{"Name": "Ethereum", "Quantity": 32},
				{"Name": "Solana", "Quantity": 52},
			},
			"Currencies": []string{"USD", "EUR", "GBP", "JPY", "CNY"},
		})

	})

	http.HandleFunc("/convertor", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/convertor.html")

		tmpl.Execute(w, map[string]interface{}{
			"selectedAmount": r.FormValue("selectedAmount"),
			"Portfolio": []map[string]interface{}{
				{"Name": "Cardano", "Quantity": 12},
				{"Name": "Bitcoin", "Quantity": 22},
				{"Name": "Ethereum", "Quantity": 32},
				{"Name": "Solana", "Quantity": 52},
			},
			"Currencies": []string{"USD", "EUR", "GBP", "JPY", "CNY"},
		})

	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
