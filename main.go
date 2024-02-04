package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

func main() {

	currencies := make([]string, 0)
	currenciesResp, _ := http.Get("https://api.coingecko.com/api/v3/simple/supported_vs_currencies")
	json.NewDecoder(currenciesResp.Body).Decode(&currencies)

	coins := make([]map[string]interface{}, 0)
	coinsResp, _ := http.Get("https://api.coingecko.com/api/v3/coins/list")
	json.NewDecoder(coinsResp.Body).Decode(&coins)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/home.html", "templates/convertor.html")

		tmpl.Execute(w, map[string]interface{}{
			"coins":      coins,
			"currencies": currencies,
		})

	})

	http.HandleFunc("/convertor", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/convertor.html")

		selectedCoin := r.FormValue("selectedCoin")
		selectedCurrency := r.FormValue("selectedCurrency")
		selectedAmount, _ := strconv.ParseFloat(r.FormValue("selectedAmount"), 64)

		data := make(map[string]map[string]float64)
		resp, _ := http.Get("https://api.coingecko.com/api/v3" + "/simple/price" + "?ids=" + selectedCoin + "&vs_currencies=" + selectedCurrency)
		json.NewDecoder(resp.Body).Decode(&data)

		tmpl.Execute(w, map[string]interface{}{
			"coins":                coins,
			"currencies":           currencies,
			"selectedAmount":       selectedAmount,
			"selectedCoin":         selectedCoin,
			"selectedCurrency":     selectedCurrency,
			"calculatedConversion": data[selectedCoin][selectedCurrency] * selectedAmount,
		})

	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
