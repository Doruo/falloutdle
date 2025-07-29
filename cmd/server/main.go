package main

import (
	"log"
	"net/http"
	"os"

	"github.com/doruo/falloutdle/cmd/server/handler"
	"github.com/doruo/falloutdle/cmd/server/handler/routes"
	"github.com/doruo/falloutdle/internal/character"
	"github.com/doruo/falloutdle/internal/database"
	"github.com/doruo/falloutdle/internal/game"
)

func main() {

	// Database connection
	db := database.GetInstance()

	// Character
	charRepo := character.NewCharacterRepository(db)
	charService := character.NewCharacterService(charRepo)

	// Game
	gameRepo := game.NewgameRepository(db)
	gameService := game.NewGameService(charService, gameRepo)
	handler := handler.NewGameHandler(gameService)

	// Creates a new game if nil
	handler.GetGameService().GetCurrentGame()

	// Server and routes setup
	mux := http.NewServeMux()
	routes.SetupRoutes(mux, handler)

	// Port listening
	host := os.Getenv("HOST")
	port := ":" + os.Getenv("PORT")
	log.Print("Server listening on http://localhost", port)
	log.Print("Website URL: https://", host)
	log.Fatal(http.ListenAndServe(port, mux))
}
