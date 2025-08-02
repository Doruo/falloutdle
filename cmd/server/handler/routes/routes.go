package routes

import (
	"net/http"

	"github.com/doruo/falloutdle/cmd/server/handler"
)

// Setups static handler api routes.
func SetupStaticRoutes(mux *http.ServeMux, h *handler.StaticHandler) {
	// GET
	const routeGetStaticHome = "/"
	mux.HandleFunc(routeGetStaticHome, h.HandleGetHome)
}

// Setups game handler api routes.
func SetupGameRoutes(mux *http.ServeMux, h *handler.GameHandler) {
	// GET
	const routeGetGames = "/api/games"
	const routeGetGameToday = "/api/games/today"
	const routeGetGameCharacterToday = "/api/characters/today"

	mux.HandleFunc(routeGetGames, h.HandleGetGames)
	mux.HandleFunc(routeGetGameToday, h.HandleGetGameToday)
	mux.HandleFunc(routeGetGameCharacterToday, h.HandleGetGameTodayCharacter)

	// POST
	const routePostGuess = "/api/games/guess"

	mux.HandleFunc(routePostGuess, h.HandlePostGuess)
}

// Setups character handler api routes.
func SetupCharacterRoutes(mux *http.ServeMux, h *handler.CharacterHandler) {
	// GET
	const routeGetCharacters = "/api/characters"
	const routeGetCharacterRandom = "/api/characters/random"

	mux.HandleFunc(routeGetCharacters, h.HandleGetCharacters)
	mux.HandleFunc(routeGetCharacterRandom, h.HandleGetCharacterRandom)
}
