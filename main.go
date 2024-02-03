package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// return html code
		fmt.Fprintf(w, "<h1>Hello, World</hf>")

	})

	fmt.Println("Starting server at port 80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
