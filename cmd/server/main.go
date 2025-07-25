package main

import (
	"log"
	"net/http"
	"os"

	"github.com/doruo/falloutdle/cmd/server/handler"
)

func main() {

	handler := handler.NewGameHandler()

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HandleGetHome)
	mux.HandleFunc("/character", handler.HandleGetCharacter)
	mux.HandleFunc("/random", handler.HandleGetRandomCharacter)

	host := os.Getenv("HOST")
	port := ":" + os.Getenv("PORT")
	log.Print("Server listening on http://", host, port)
	log.Fatal(http.ListenAndServe(port, mux))
}
