package routes

import (
	"net/http"

	"github.com/doruo/falloutdle/cmd/server/handler"
)

// home
const routeGetHome = "/"

// Game routes
// GET
const routeGetGames = "/api/games"
const routeGetGameToday = "/api/games/today"

//const rougeGetGame = "/api/games/:id"

// POST
const routePostGuess = "/api/guess"

// Character routes
// GET
const routeGetCharacters = "/api/characters"

// const routeGetCharacter = "/api/characters/:id"
const routeGetCharacterToday = "/api/characters/today"
const routeGetCharacterRandom = "/api/characters/random"

func SetupGameRoutes(mux *http.ServeMux, h *handler.GameHandler) {
	// HOME
	mux.HandleFunc(routeGetHome, h.HandleGetHome)

	// GET
	mux.HandleFunc(routeGetGames, h.HandleGetGames)
	mux.HandleFunc(routeGetGameToday, h.HandleGetGameToday)
	mux.HandleFunc(routeGetCharacterToday, h.HandleGetTodayCharacter)

	// POST
	mux.HandleFunc(routePostGuess, h.HandlePostGuess)
}

func SetupCharacterRoutes(mux *http.ServeMux, h *handler.CharacterHandler) {
	// GET
	mux.HandleFunc(routeGetCharacters, h.HandleGetCharacters)
	mux.HandleFunc(routeGetCharacterRandom, h.HandleGetCharacterRandom)
}
