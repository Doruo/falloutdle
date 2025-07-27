package main

import (
	"log"
	"net/http"
	"os"

	"github.com/doruo/falloutdle/cmd/server/handler/routes"
)

func main() {

	mux := http.NewServeMux()
	routes.SetupRoutes(mux)

	host := os.Getenv("HOST")
	port := ":" + os.Getenv("PORT")
	log.Print("Server listening on http://localhost", port)
	log.Print("Website URL: ", host)
	log.Fatal(http.ListenAndServe(port, mux))
}
