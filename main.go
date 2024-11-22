package main

import (
	"net/http"

	"HTTP-SERVER/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}