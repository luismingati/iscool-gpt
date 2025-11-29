package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"iscool-gpt/internal/gemini"
)

type PromptRequest struct {
	Prompt string `json:"prompt"`
}

type PromptResponse struct {
	Response string `json:"response"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type PromptHandler struct {
	geminiClient *gemini.Client
}

func NewPromptHandler(geminiClient *gemini.Client) *PromptHandler {
	return &PromptHandler{
		geminiClient: geminiClient,
	}
}

func (h *PromptHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req PromptRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Invalid JSON: " + err.Error(),
		}); err != nil {
			log.Printf("Error encoding error response: %v", err)
		}
		return
	}

	if req.Prompt == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Prompt cannot be empty",
		}); err != nil {
			log.Printf("Error encoding error response: %v", err)
		}
		return
	}

	response, err := h.geminiClient.GenerateResponse(r.Context(), req.Prompt)
	if err != nil {
		log.Printf("Error generating response: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(ErrorResponse{
			Error: "Failed to generate response: " + err.Error(),
		}); err != nil {
			log.Printf("Error encoding error response: %v", err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(PromptResponse{
		Response: response,
	}); err != nil {
		log.Printf("Error encoding success response: %v", err)
	}
}
