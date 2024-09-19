package main

import (
	"fmt"
	"net/http"

	handler "github.com/0jk6/triemap/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handler.HomeHandler)

	mux.HandleFunc("GET /get/", handler.GetHandler)
	mux.HandleFunc("POST /store", handler.StoreHandler)

	mux.HandleFunc("GET /prefix/", handler.PrefixSearchHandler)
	mux.HandleFunc("GET /suffix/", handler.SuffixSearchHandler)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
