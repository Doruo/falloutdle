package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/doruo/falloutdle/pkg/time"
)

// JSON response handler format
type Response struct {
	Success bool   `json:"success"`
	Data    []any  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
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

	fmt.Println(time.Today(), "API - HTTP error", httpStatus, ":", message)
	sendJSONResponse(writer, Response{
		Success: false,
		Data:    nil,
		Error:   message,
	})
}

func sendReponse(writer http.ResponseWriter, response Response) {
	if err := json.NewEncoder(writer).Encode(&response); err != nil {
		http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
	}
}
