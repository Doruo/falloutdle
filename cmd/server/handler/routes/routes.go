package routes

import (
	"net/http"

	"github.com/doruo/falloutdle/cmd/server/handler"
)

func SetupRoutes(mux *http.ServeMux) {
	handler := handler.NewGameHandler()

	mux.HandleFunc("/", handler.HandleGetHome)
	mux.HandleFunc("/api/today", handler.HandleGetTodayCharacter)
	mux.HandleFunc("/api/random", handler.HandleGetRandomCharacter)
	mux.HandleFunc("/api/guess", handler.HandlePostGuessCharacter)
}
