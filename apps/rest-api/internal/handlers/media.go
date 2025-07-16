package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"rest-api/internal/models"
)

// SaveExerciseMedia godoc
// @Summary Save exercise media
// @Description Save media (images, videos) for an exercise
// @Tags media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.ExerciseMediaRequest true "Exercise media data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /api/exercise/media [post]
func (h *Handlers) SaveExerciseMedia(w http.ResponseWriter, r *http.Request) {
	var req models.ExerciseMediaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request format")
		return
	}

	if err := h.MediaService.SaveExerciseMedia(r.Context(), &req); err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{
		"message": "Exercise media saved successfully",
	})
}

// GetExerciseMedia godoc
// @Summary Get exercise media
// @Description Get all media for a specific exercise
// @Tags media
// @Produce json
// @Security BearerAuth
// @Param exercise_id path string true "Exercise ID"
// @Success 200 {array} models.ExerciseMedia
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /api/exercise/{exercise_id}/media [get]
func (h *Handlers) GetExerciseMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	exerciseID := vars["exercise_id"]

	media, err := h.MediaService.GetExerciseMedia(r.Context(), exerciseID)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, media)
}

// DeleteExerciseMedia godoc
// @Summary Delete exercise media
// @Description Delete a specific media item
// @Tags media
// @Produce json
// @Security BearerAuth
// @Param media_id path string true "Media ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /api/exercise/media/{media_id} [delete]
func (h *Handlers) DeleteExerciseMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mediaID := vars["media_id"]

	if err := h.MediaService.DeleteExerciseMedia(r.Context(), mediaID); err != nil {
		handleServiceError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Exercise media deleted successfully",
	})
}
