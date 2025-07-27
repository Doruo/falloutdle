package routes

import (
	"net/http"

	"github.com/doruo/falloutdle/cmd/server/handler"
	"github.com/doruo/falloutdle/internal/character"
	"github.com/doruo/falloutdle/internal/database"
	"github.com/doruo/falloutdle/internal/game"
)

func SetupRoutes(mux *http.ServeMux) {

	db := database.GetInstance()

	charRepo := character.NewCharacterRepository(db)
	charService := character.NewCharacterService(charRepo)

	gameRepo := game.NewgameRepository(db)
	gameService := game.NewGameService(charService, gameRepo)
	handler := handler.NewGameHandler(gameService)

	// Creates a new game if nil
	handler.GetGameService().GetCurrentCharacter()

	mux.HandleFunc("/", handler.HandleGetHome)
	mux.HandleFunc("/api/today", handler.HandleGetTodayCharacter)
	mux.HandleFunc("/api/random", handler.HandleGetRandomCharacter)
	mux.HandleFunc("/api/guess", handler.HandlePostGuessCharacter)
}
