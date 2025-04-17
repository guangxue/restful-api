package main

import (
	"net/http"

	"github.com/guangxue/restful-api/internal/app/hongfa"
)

func main() {
	router := http.NewServeMux()

	// Hongfa-IMS Api
	router.HandleFunc("GET /hongfa/{$}", hongfa.Hongfa)
	router.HandleFunc("GET /hongfa/items/{$}", hongfa.Items)

	http.ListenAndServe(":6666", router)
}
