package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type StaticHandler struct{}

func NewStaticHandler() *StaticHandler {
	return &StaticHandler{}
}

func (h *StaticHandler) HandleGetHome(w http.ResponseWriter, r *http.Request) {

	fmt.Println("API - handling GET request: home page")

	// Verify correct http method
	if !isGetMethod(r) {
		sendErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	frontendPath := os.Getenv("FRONTEND_PATH")

	// Serve index html file
	if r.URL.Path == "/" {
		http.ServeFile(w, r, filepath.Join(frontendPath, "index.html"))
		return
	}

	// Serve other files
	filePath := filepath.Join(frontendPath, r.URL.Path)
	http.ServeFile(w, r, filePath)
}
