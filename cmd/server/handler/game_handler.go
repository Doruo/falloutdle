package handler

import (
	"encoding/json"
	"net/http"

	"github.com/doruo/falloutdle/external/game"
)

type GameHandler struct {
	gameService *game.GameService
}

func NewGameHandler(gs *game.GameService) *GameHandler {
	return &GameHandler{
		gameService: gs,
	}
}

// /----- HTTP GET FUNCTIONS -----/

// HandleDefault
func (handler *GameHandler) HandleDefault(writer http.ResponseWriter, request *http.Request) {

	// Verify correct http method
	if !isMethod(request.Method, http.MethodGet) {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sendResponse(writer, "Hello")
}

// HandleGetCharacter returns today guess character.
func (handler *GameHandler) HandleGetCharacter(writer http.ResponseWriter, request *http.Request) {

	// Verify correct http method.
	if !isMethod(request.Method, http.MethodGet) {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(request.Body).Decode(&request); err != nil {
		http.Error(writer, "Invalid JSON", http.StatusBadRequest)
		return
	}

	character, error := handler.gameService.GetTodayCharacter()

	if error != nil {
		http.Error(writer, "Error while getting character", http.StatusInternalServerError)
		return
	}

	sendResponse(writer, character.Name)
}

// sendResponse sends response with content in json format.
func sendResponse(w http.ResponseWriter, content any) {

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(content); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// isMethod verify correct HTTP method.
func isMethod(requestMethod string, validMethod string) bool {
	return requestMethod == validMethod
}
