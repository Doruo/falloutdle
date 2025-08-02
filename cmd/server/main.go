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

	// Character service setup
	characterRepo := character.NewCharacterRepository(db)
	characterService := character.NewCharacterService(characterRepo)

	// Game service setup
	gameRepo := game.NewgameRepository(db)
	gameService := game.NewGameService(characterService, gameRepo)
	gameService.GetGameCurrent() // Creates a new game if nil

	// Handlers setup
	staticHandler := handler.NewStaticHandler()
	characterHandler := handler.NewCharacterHandler(characterService)
	gameHandler := handler.NewGameHandler(gameService)

	// Frontend page handling
	frontendPath := os.Getenv("FRONTEND_PATH")

	// Router server setup
	mux := http.NewServeMux()

	// Api routes
	routes.SetupGameRoutes(mux, gameHandler)
	routes.SetupCharacterRoutes(mux, characterHandler)

	// Static routes (home page, assets)
	routes.SetupStaticRoutes(mux, staticHandler)

	// Logs
	host := os.Getenv("HOST")
	port := ":" + os.Getenv("PORT")

	log.Print("Frontend path:", frontendPath)
	log.Print("Server listening on http://localhost", port)
	log.Print("Website URL: https://", host)

	// Port listening
	log.Fatal(http.ListenAndServe(port, mux))
}
