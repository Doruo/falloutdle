package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/doruo/falloutdle/cmd/server/handler"
	"github.com/doruo/falloutdle/cmd/server/handler/routes"
	"github.com/doruo/falloutdle/internal/character"
	"github.com/doruo/falloutdle/internal/database"
	"github.com/doruo/falloutdle/internal/game"
)

func displayTitle() {
	file, _ := os.ReadFile("../../title.txt")
	fmt.Println()
	fmt.Println(string(file))
	fmt.Println()
}

func main() {
	displayTitle()

	// Database connection
	db := database.GetInstance()

	// Character
	characterRepo := character.NewCharacterRepository(db)
	characterService := character.NewCharacterService(characterRepo)

	// Game
	gameRepo := game.NewGameRepository(db)
	gameService := game.NewGameService(characterService, gameRepo)
	gameService.GetGameCurrent() // Creates a new game if nil

	// Handlers
	characterHandler := handler.NewCharacterHandler(characterService)
	gameHandler := handler.NewGameHandler(gameService)

	// Server and routes setup
	mux := http.NewServeMux()
	routes.SetupGameRoutes(mux, gameHandler)
	routes.SetupCharacterRoutes(mux, characterHandler)

	listen(mux)
}

func listen(mux *http.ServeMux) {
	// Port listening
	host := os.Getenv("HOST")
	port := ":" + os.Getenv("PORT")
	log.Print("Server listening on http://localhost", port)
	log.Print("Website URL: https://", host)
	log.Fatal(http.ListenAndServe(port, mux))
}
