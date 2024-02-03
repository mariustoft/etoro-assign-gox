package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Starting server at port 80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
