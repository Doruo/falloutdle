package handler

import (
	"fmt"
	"net/http"

	"github.com/doruo/falloutdle/internal/character"
)

type CharacterHandler struct {
	characterService *character.Service
}

func NewCharacterHandler(ch *character.Service) *CharacterHandler {
	return &CharacterHandler{
		characterService: ch,
	}
}

// HandleGetTodayCharacter returns today guess character.
func (h *CharacterHandler) HandleGetCharacters(w http.ResponseWriter, r *http.Request) {

	fmt.Println("API - handling GET request: all characters")

	// Verify correct http method
	if !isGetMethod(r) {
		sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	characters, error := h.characterService.GetCharacters()

	if error != nil {
		sendErrorResponse(w, "Error while getting character", http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, Response{
		Success:    true,
		Data:       characters,
		DataLength: len(characters),
	})
}
