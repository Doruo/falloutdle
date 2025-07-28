package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/doruo/falloutdle/internal/game"
	"github.com/doruo/falloutdle/pkg/time"
)

type GameHandler struct {
	gameService *game.GameService
}

func NewGameHandler(gs *game.GameService) *GameHandler {
	return &GameHandler{
		gameService: gs,
	}
}

// /----- HTTP GET -----/

// HandleGetHome
func (handler *GameHandler) HandleGetHome(w http.ResponseWriter, r *http.Request) {

	fmt.Println(time.Today(), "API - handling GET request: home page ...")

	// Verify correct http method
	if !isGetMethod(r) {
		sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	const url = "./index.html"
	content, err := os.ReadFile(url)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	sendHTMLResponse(w, content)
}

// HandleGetTodayCharacter returns today guess character.
func (handler *GameHandler) HandleGetTodayCharacter(w http.ResponseWriter, r *http.Request) {

	fmt.Println("API - handling GET request: today character")

	// Verify correct http method
	if !isGetMethod(r) {
		sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	character, error := handler.gameService.GetCurrentCharacter()

	if error != nil {
		sendErrorResponse(w, "Error while getting character", http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, Response{
		Success: true,
		Data:    []any{character},
	})
}

// HandleGetRandomCharacter returns random character from fallout games.
func (handler *GameHandler) HandleGetRandomCharacter(w http.ResponseWriter, r *http.Request) {

	fmt.Println(time.Today(), "API - handling GET request: random character")

	// Verify correct http method
	if !isGetMethod(r) {
		sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	character, error := handler.gameService.GetRandomCharacter()

	if error != nil {
		sendErrorResponse(w, "Error while getting character", http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, Response{
		Success: true,
		Data:    []any{character},
	})
}

// /----- HTTP POST -----/

// HandlePostGuessCharacter receive and process character guess attempt.
func (handler *GameHandler) HandlePostGuessCharacter(w http.ResponseWriter, r *http.Request) {

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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Parse result result
	var result map[string]string
	err = json.Unmarshal(body, &result)

	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := result["character_name"]
	fmt.Println("Guess value:", name)
	isGuessed, err := handler.gameService.ProcessGuess(name)

	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Is correct:", isGuessed)

	sendJSONResponse(w, Response{
		Success: true,
		Data:    []any{GuessResponse{IsGuessed: isGuessed}},
	})
}

// /----- GET FUNCTIONS -----/

// HandlePostGuessCharacter receive and process character guess attempt.
func (handler *GameHandler) GetGameService() *game.GameService {
	return handler.gameService
}
