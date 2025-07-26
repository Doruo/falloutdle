package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/doruo/falloutdle/external/game"
)

type GameHandler struct {
	gameService *game.GameService
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		gameService: game.GetInstance(),
	}
}

// /----- HTTP GET FUNCTIONS -----/

// HandleGetHome
func (handler *GameHandler) HandleGetHome(writer http.ResponseWriter, request *http.Request) {

	// Verify correct http method
	if !isMethod(request.Method, http.MethodGet) {
		sendResponseError(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	const url = "./index.html"
	content, err := os.ReadFile(url)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	sendResponseHTML(writer, content)
}

// HandleGetCharacter returns today guess character.
func (handler *GameHandler) HandleGetCharacter(writer http.ResponseWriter, request *http.Request) {

	// Verify correct http method
	if !isMethod(request.Method, http.MethodGet) {
		sendResponseError(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	character, error := handler.gameService.GetCurrentCharacter()

	if error != nil {
		sendResponseError(writer, "Error while getting character", http.StatusInternalServerError)
		return
	}

	response := Response{
		Success: true,
		Data:    character.Name,
	}

	sendResponseJSON(writer, response)
}

// HandleGetCharacter returns today guess character.
func (handler *GameHandler) HandleGetRandomCharacter(writer http.ResponseWriter, request *http.Request) {

	// Verify correct http method
	if !isMethod(request.Method, http.MethodGet) {
		sendResponseError(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	character, error := handler.gameService.GetRandomCharacter()

	if error != nil {
		sendResponseError(writer, "Error while getting character", http.StatusInternalServerError)
		return
	}

	response := Response{
		Success: true,
		Data:    character.Name,
	}

	sendResponseJSON(writer, response)
}

// /----- SEND RESPONSE METHODS -----/

// sendResponseJSON sends response with content in json format.
func sendResponseJSON(writer http.ResponseWriter, response Response) {
	writer.Header().Set("Content-Type", "application/json")
	sendReponse(writer, response)
}

func sendResponseHTML(writer http.ResponseWriter, content []byte) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.Write(content)
}

// sendResponseJSON sends response with content in json format.
func sendResponseError(writer http.ResponseWriter, message string, httpStatus int) {
	response := Response{
		Success: false,
		Error:   message,
	}
	writer.WriteHeader(httpStatus)
	sendReponse(writer, response)
}

func sendReponse(writer http.ResponseWriter, response Response) {
	if err := json.NewEncoder(writer).Encode(&response); err != nil {
		http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
	}
}

// /----- UTILITY METHODS -----/

// isMethod verify correct HTTP method.
func isMethod(requestMethod string, validMethod string) bool {
	return requestMethod == validMethod
}
