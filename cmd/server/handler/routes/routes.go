package routes

import (
	"net/http"

	"github.com/doruo/falloutdle/cmd/server/handler"
)

func SetupStaticRoutes(mux *http.ServeMux, h *handler.StaticHandler) {
	// GET
	const routeGetStaticHome = "/"
	mux.HandleFunc(routeGetStaticHome, h.HandleGetHome)
}

func SetupGameRoutes(mux *http.ServeMux, h *handler.GameHandler) {
	// GET
	const routeGetGames = "/api/games"
	const routeGetGameToday = "/api/games/today"

	mux.HandleFunc(routeGetGames, h.HandleGetGames)
	mux.HandleFunc(routeGetGameToday, h.HandleGetGameToday)

	// POST
	const routePostGuess = "/api/guess"

	mux.HandleFunc(routePostGuess, h.HandlePostGuess)
}

func SetupCharacterRoutes(mux *http.ServeMux, h *handler.CharacterHandler) {
	// GET
	const routeGetCharacters = "/api/characters"
	const routeGetCharacterRandom = "/api/characters/random"

	mux.HandleFunc(routeGetCharacters, h.HandleGetCharacters)
	mux.HandleFunc(routeGetCharacterRandom, h.HandleGetCharacterRandom)
}
