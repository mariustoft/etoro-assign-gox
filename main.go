package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
)

var portfolio = []map[string]interface{}{
	{"Name": "cardano", "Quantity": 12},
	{"Name": "bitcoin", "Quantity": 22},
	{"Name": "ethereum", "Quantity": 32},
	{"Name": "eolana", "Quantity": 52},
}

var currencies = []string{"usd", "eur", "gbp", "jpy", "cny"}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/home.html", "templates/convertor.html")

		tmpl.Execute(w, map[string]interface{}{
			"currencies": currencies,
			"portfolio":  portfolio,
		})

	})

	http.HandleFunc("/convertor", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/convertor.html")

		selectedCoin := r.FormValue("selectedCoin")
		selectedCurrency := r.FormValue("selectedCurrency")
		selectedAmount, _ := strconv.ParseFloat(r.FormValue("selectedAmount"), 64)

		url := "https://api.coingecko.com/api/v3" + "/simple/price"
		url += "?ids=" + selectedCoin + "&vs_currencies=" + selectedCurrency + "&amount=" + strconv.FormatFloat(selectedAmount, 'f', -1, 64)

		resp, _ := http.Get(url)
		body, _ := io.ReadAll(resp.Body)

		data := make(map[string]map[string]float64)
		json.Unmarshal([]byte(body), &data)

		tmpl.Execute(w, map[string]interface{}{
			"selectedAmount":       selectedAmount,
			"selectedCoin":         selectedCoin,
			"selectedCurrency":     selectedCurrency,
			"calculatedConversion": data[selectedCoin][selectedCurrency] * selectedAmount,
			"currencies":           currencies,
			"portfolio":            portfolio,
		})

	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
