package main

import (
	"html/template"
	"net/http"
	"os"
)

func main() {
	tmpl, _ := template.ParseFiles("home.html", "header.html")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl.Execute(w, map[string]interface{}{
			"Header": r.FormValue("new-header"),
			"Portfolio": []map[string]interface{}{
				{"Name": "BTC", "Price": 10000.0},
				{"Name": "ETH", "Price": 200.0},
				{"Name": "XRP", "Price": 0.5},
			},
		})

	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
