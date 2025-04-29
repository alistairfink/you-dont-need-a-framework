package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type (
	exampleRequest struct {
		Word string `json:"word"`
	}
	exampleResponse struct {
		Message string `json:"message"`
	}
)

type ExampleHandler struct {
}

func NewExampleHandler() *ExampleHandler {
	return &ExampleHandler{}
}

func (h *ExampleHandler) initializeRoutes() []route {
	return []route{
		{
			method:  http.MethodPost,
			path:    "/message",
			handler: h.createMessage,
		},
	}
}

func (h *ExampleHandler) createMessage(w http.ResponseWriter, r *http.Request) {
	logger, _ := r.Context().Value("logger").(*slog.Logger)
	logger.Info("test")
	var params exampleRequest
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		logger.Error("Error decoding params", slog.Any("error", err))
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	response := exampleResponse{
		Message: "Hello " + params.Word,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Error encoding response", slog.Any("error", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
