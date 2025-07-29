package routes

import (
	"net/http"

	"github.com/doruo/falloutdle/cmd/server/handler"
)

// GET routes
const routeGetHome = "/"
const routeGetTodayCharacter = "/api/today"
const routeGetRandomCharacter = "/api/random"

// POST routes
const routePostGuess = "/api/guess"

func SetupRoutes(mux *http.ServeMux, handler *handler.GameHandler) {

	// GET routes
	mux.HandleFunc(routeGetHome, handler.HandleGetHome)
	mux.HandleFunc(routeGetTodayCharacter, handler.HandleGetTodayCharacter)
	mux.HandleFunc(routeGetRandomCharacter, handler.HandleGetRandomCharacter)

	// POST routes
	mux.HandleFunc(routePostGuess, handler.HandlePostGuessCharacter)
}
