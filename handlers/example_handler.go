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
			method:          http.MethodPost,
			path:            "/with_response_code",
			handler:         h.createMessage,
			logResponseCode: true,
		},
		{
			method:          http.MethodPost,
			path:            "/without_response_code",
			handler:         h.createMessage,
			logResponseCode: false,
		},
		{
			method:          http.MethodPost,
			path:            "/with_response_code_error",
			handler:         h.createMessageError,
			logResponseCode: true,
		},
		{
			method:  http.MethodPost,
			path:    "/without_response_code_error",
			handler: h.createMessageError,
			// logResponseCode defaults to false
		},
	}
}

func (h *ExampleHandler) createMessage(w http.ResponseWriter, r *http.Request) {
	logger, _ := r.Context().Value("logger").(*slog.Logger)
	logger.Info("without error")
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

func (h *ExampleHandler) createMessageError(w http.ResponseWriter, r *http.Request) {
	logger, _ := r.Context().Value("logger").(*slog.Logger)
	logger.Info("with error")
	http.Error(w, "Invalid request", http.StatusBadRequest)
}
