package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/doruo/falloutdle/external/game"
	"github.com/doruo/falloutdle/pkg/time"
)

type GameHandler struct {
	gameService *game.GameService
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		gameService: game.GetServiceInstance(),
	}
}

// /----- HTTP GET -----/

// HandleGetHome
func (handler *GameHandler) HandleGetHome(writer http.ResponseWriter, request *http.Request) {

	fmt.Println(time.Today(), "API - handling GET request: home page ...")

	// Verify correct http method
	if !isMethod(request.Method, http.MethodGet) {
		sendErrorResponse(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	const url = "./index.html"
	content, err := os.ReadFile(url)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	sendHTMLResponse(writer, content)
}

// HandleGetTodayCharacter returns today guess character.
func (handler *GameHandler) HandleGetTodayCharacter(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("API - handling GET request: today character")

	// Verify correct http method
	if !isMethod(request.Method, http.MethodGet) {
		sendErrorResponse(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	character, error := handler.gameService.GetCurrentCharacter()

	if error != nil {
		sendErrorResponse(writer, "Error while getting character", http.StatusInternalServerError)
		return
	}

	sendJSONResponse(writer, Response{
		Success: true,
		Data:    []any{character},
	})
}

// HandleGetRandomCharacter returns random character from fallout games.
func (handler *GameHandler) HandleGetRandomCharacter(writer http.ResponseWriter, request *http.Request) {

	fmt.Println(time.Today(), "API - handling GET request: random character")

	// Verify correct http method
	if !isMethod(request.Method, http.MethodGet) {
		sendErrorResponse(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	character, error := handler.gameService.GetRandomCharacter()

	if error != nil {
		sendErrorResponse(writer, "Error while getting character", http.StatusInternalServerError)
		return
	}

	sendJSONResponse(writer, Response{
		Success: true,
		Data:    []any{character},
	})
}

// /----- HTTP POST -----/

func (handler *GameHandler) HandlePostGuessCharacter(writer http.ResponseWriter, request *http.Request) {

	fmt.Println(time.Today(), "API - handling POST request: guess character")

	// Verify correct http method
	if !isMethod(request.Method, http.MethodPost) {
		sendErrorResponse(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("guess character:", request.Body)
}

// /----- UTILITY METHODS -----/

// isMethod verify correct HTTP method.
func isMethod(requestMethod string, validMethod string) bool {
	return requestMethod == validMethod
}
