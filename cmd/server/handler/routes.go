package handler

import (
	"net/http"
)

func SetupRoutes(mux *http.ServeMux, handler *GameHandler) {
	mux.HandleFunc("/", handler.HandleGetHome)
	mux.HandleFunc("/api/character", handler.HandleGetCharacter)
	mux.HandleFunc("/api/random", handler.HandleGetRandomCharacter)
}
