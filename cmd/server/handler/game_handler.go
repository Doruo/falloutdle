package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/doruo/falloutdle/internal/game"
	"github.com/doruo/falloutdle/pkg/time"
)

type GameHandler struct {
	gameService *game.Service
}

func NewGameHandler(gs *game.Service) *GameHandler {
	return &GameHandler{
		gameService: gs,
	}
}

// /----- HTTP GET -----/

func (h *GameHandler) HandleGetGames(w http.ResponseWriter, r *http.Request) {

	fmt.Println("API - handling GET request: all games")

	// Verify correct http method
	if !isGetMethod(r) {
		sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	games, error := h.gameService.GetGames()

	if error != nil {
		sendErrorResponse(w, "Error while getting games", http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, Response{
		Success:    true,
		Data:       games,
		DataLength: len(games),
	})
}

func (h *GameHandler) HandleGetGameToday(w http.ResponseWriter, r *http.Request) {

	fmt.Println("API - handling GET request: today game")

	// Verify correct http method
	if !isGetMethod(r) {
		sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	game, error := h.gameService.GetGameCurrent()

	if error != nil {
		sendErrorResponse(w, "Error while getting character: "+error.Error(), http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, Response{
		Success: true,
		Data:    game,
	})
}

// HandleGetTodayCharacter returns today guess character.
func (h *GameHandler) HandleGetTodayCharacter(w http.ResponseWriter, r *http.Request) {

	fmt.Println("API - handling GET request: today character")

	// Verify correct http method
	if !isGetMethod(r) {
		sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	character, error := h.gameService.GetGameCurrentCharacter()

	if error != nil {
		sendErrorResponse(w, "Error while getting character: "+error.Error(), http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, Response{
		Success: true,
		Data:    character,
	})
}

// /----- HTTP POST -----/

// HandlePostGuess receive and process character guess attempt.
func (h *GameHandler) HandlePostGuess(w http.ResponseWriter, r *http.Request) {

	fmt.Println(time.Today(), "API - handling POST request: guess character")

	// Verify correct http method
	if !isPostMethod(r) {
		sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Verify correct content-type
	if !isContentTypeJSON(&r.Header) {
		sendErrorResponse(w, "Content-Type must be application/json", http.StatusBadRequest)
		return
	}

	// Read body
	defer r.Body.Close()
	body, error := io.ReadAll(r.Body)
	if error != nil {
		sendErrorResponse(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Parse result result
	var result map[string]string
	error = json.Unmarshal(body, &result)

	if error != nil {
		sendErrorResponse(w, error.Error(), http.StatusBadRequest)
		return
	}

	name := result["character_name"]
	fmt.Println("Guess value:", name)
	isGuessed, error := h.gameService.ProcessGuess(name)

	if error != nil {
		sendErrorResponse(w, error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Is correct:", isGuessed)

	sendJSONResponse(w, Response{
		Success: true,
		Data:    GuessResponse{IsGuessed: isGuessed},
	})
}

// /----- GET FUNCTIONS -----/

// HandlePostGuessCharacter receive and process character guess attempt.
func (h *GameHandler) GetGameService() *game.Service {
	return h.gameService
}
