package main

import (
	"log"
	"net/http"
	"os"

	"github.com/doruo/falloutdle/cmd/server/handler"
)

func main() {

	mux := http.NewServeMux()
	handler.SetupRoutes(mux)

	host := os.Getenv("HOST")
	port := ":" + os.Getenv("PORT")
	log.Print("Server listening on http://", host, port)
	log.Fatal(http.ListenAndServe(port, mux))
}
