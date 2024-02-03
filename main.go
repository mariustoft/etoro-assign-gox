package main

import (
	"fmt"
	"net/http"
	// "os"
	// "todox/internal/app"
	// "todox/internal/app/config"
	// "github.com/leapkit/core/server"
)

// func main() {
// 	s := server.New(
// 		"Todox",

// 		server.WithPort(config.Port),
// 		server.WithHost(config.Host),
// 	)

// 	// Application services
// 	if err := app.AddServices(s); err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	// Application routes
// 	if err := app.AddRoutes(s); err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	if err := s.Start(); err != nil {
// 		fmt.Println(err)
// 	}
// }

func main() {
	// simple hellow work h1 server for webpage with native go
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ddd?d")
	})

	http.ListenAndServe(":8080", nil)

}
