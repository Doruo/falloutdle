package main

import (
	"log"
	"net/http"
	"os"

	"github.com/doruo/falloutdle/cmd/server/handler"
)

func main() {

	handler := handler.GetInstance()
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HandleDefault)

	port := os.Getenv("PORT")
	portStr := ":" + port

	log.Println("Server listening on http://localhost:", port)
	log.Fatal(http.ListenAndServe(portStr, mux))
}
