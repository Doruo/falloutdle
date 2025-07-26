package handler

import (
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	handler := NewGameHandler()
	mux.HandleFunc("/", handler.HandleGetHome)
	mux.HandleFunc("/api/today", handler.HandleGetCharacter)
	mux.HandleFunc("/api/random", handler.HandleGetRandomCharacter)
}
