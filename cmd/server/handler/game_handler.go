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
		gameService: game.GetServiceInstance(),
	}
}

// /----- HTTP GET FUNCTIONS -----/

// HandleGetHome
func (handler *GameHandler) HandleGetHome(writer http.ResponseWriter, request *http.Request) {

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

	fmt.Println("API: today character request...")

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

	response := Response{
		Success: true,
		Data:    []any{character},
	}

	sendJSONResponse(writer, response)
}

// HandleGetRandomCharacter returns random character from fallout games.
func (handler *GameHandler) HandleGetRandomCharacter(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("API: Handling random character request...")

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

// /----- SEND RESPONSE METHODS -----/

// sendJSONResponse sends response with content in JSON format.
func sendJSONResponse(writer http.ResponseWriter, response Response) {
	writer.Header().Set("Content-Type", "application/json")
	sendReponse(writer, response)
}

// sendHTMLResponse sends response with content in HTML format.
func sendHTMLResponse(writer http.ResponseWriter, content []byte) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.Write(content)
}

// sendErrorResponse sends response error with message and httpStatus in json format.
func sendErrorResponse(writer http.ResponseWriter, message string, httpStatus int) {

	fmt.Println("API error :", message, httpStatus)

	writer.WriteHeader(httpStatus)
	sendJSONResponse(writer, Response{
		Error: message,
	})
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
