package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/doruo/falloutdle/pkg/time"
)

// JSON request handler format
type GuessRequest struct {
	CharacterName string `json:"character_name"`
}

type GuessResponse struct {
	IsGuessed bool `json:"isGuessed"`
}

// JSON response handler format
type Response struct {
	Success bool   `json:"success"`
	Data    []any  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// /----- SEND RESPONSE METHODS -----/

// sendJSONResponse sends response with content in JSON format.
func sendJSONResponse(w http.ResponseWriter, r Response) {
	w.Header().Set("Content-Type", "application/json")
	sendReponse(w, r)
}

// sendHTMLResponse sends response with content in HTML format.
func sendHTMLResponse(w http.ResponseWriter, content []byte) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(content)
}

// sendErrorResponse sends response error with message and httpStatus in json format.
func sendErrorResponse(w http.ResponseWriter, msg string, httpStatus int) {

	fmt.Println(time.Today(), "API - HTTP error", httpStatus, ":", msg)

	sendJSONResponse(w, Response{
		Success: false,
		Data:    nil,
		Error:   msg,
	})
}

func sendReponse(w http.ResponseWriter, r Response) {
	if err := json.NewEncoder(w).Encode(r); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// /----- UTILITY METHODS -----/

// isGetMethod verify correct GET HTTP method.
func isGetMethod(r *http.Request) bool {
	return r.Method == http.MethodGet
}

// isGetMethod verify correct POST HTTP method.
func isPostMethod(r *http.Request) bool {
	return r.Method == http.MethodPost
}

// isContentTypeJSON.
func isContentTypeJSON(h *http.Header) bool {
	return h.Get("Content-Type") == "application/json"
}
