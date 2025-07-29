package routes

import (
	"net/http"

	"github.com/doruo/falloutdle/cmd/server/handler"
)

// home
const routeGetHome = "/"

// Game routes
// GET
//const routeGetGames = "/api/games"
//const rougeGetGame = "/api/games/:id"

// POST
const routePostGuess = "/api/guess"

// Character routes
// GET
const routeGetCharacters = "/api/characters"

// const routeGetCharacter = "/api/characters/:id"
const routeGetCharacterToday = "/api/characters/today"
const routeGetCharacterRandom = "/api/characters/random"

func SetupGameRoutes(mux *http.ServeMux, gh *handler.GameHandler) {
	// GET
	mux.HandleFunc(routeGetHome, gh.HandleGetHome)
	mux.HandleFunc(routeGetCharacterToday, gh.HandleGetTodayCharacter)
	mux.HandleFunc(routeGetCharacterRandom, gh.HandleGetRandomCharacter)

	// POST
	mux.HandleFunc(routePostGuess, gh.HandlePostGuessCharacter)
}

func SetupCharacterRoutes(mux *http.ServeMux, ch *handler.CharacterHandler) {
	// GET
	mux.HandleFunc(routeGetCharacters, ch.HandleGetCharacters)
}
