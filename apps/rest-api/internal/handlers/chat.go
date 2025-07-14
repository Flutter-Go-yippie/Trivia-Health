package handlers

import (
	"encoding/json"
	"net/http"

	"rest-api/internal/models"
)

func (h *Handlers) Chat(w http.ResponseWriter, r *http.Request) {
	var req models.ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request format")
		return
	}

	response, err := h.AIService.Chat(r.Context(), req.Message)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, models.ChatResponse{
		Response: response,
	})
}

func (h *Handlers) GetChatHistory(w http.ResponseWriter, r *http.Request) {
	history, err := h.AIService.GetChatHistory(r.Context())
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, models.ChatHistory{
		Messages: history,
	})
}
