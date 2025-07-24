package main

import (
	"log"
	"net/http"
	"os"

	"github.com/doruo/falloutdle/cmd/server/handler"
	"github.com/doruo/falloutdle/external/game"
	"github.com/doruo/falloutdle/internal/character"
	"github.com/doruo/falloutdle/internal/database"
)

func main() {

	db := database.NewDatabaseConnection()
	repo := character.NewCharacterRepository(db)
	characterService := character.NewCharacterService(repo)
	gameService := game.NewGameService(characterService)
	gameHandler := handler.NewGameHandler(gameService)

	mux := http.NewServeMux()

	mux.HandleFunc("/", gameHandler.HandleDefault)

	port := os.Getenv("PORT")
	portStr := ":" + port

	log.Println("Server listening on http://localhost:", port)
	log.Fatal(http.ListenAndServe(portStr, mux))
}
