package main

import (
	"WalletViewAPI/api/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const port = "8080"

func main() {
	// read the port from env var
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handler.GetWalletBalance)
	log.Printf("Starting server at port :%s\n", port)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%s", port),
		r,
	))
}
