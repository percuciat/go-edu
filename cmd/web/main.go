package main

import (
	"goProject/internal/web/router"
	"log"
	"net/http"
)

func main() {
	r := router.Initialize()

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
