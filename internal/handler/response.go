package handler

import (
	"encoding/json"
	"net/http"
)

// Response - Represent http response message
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// NewResponse - Init and throw a http response message
func (h *Handler) NewResponse(w http.ResponseWriter, statusCode int, message string) {
	resp := Response{
		Status:  statusCode,
		Message: message,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		h.logger.WithError(err).Error("failed to marshal response")
		return
	}

	w.WriteHeader(statusCode)
	_, _ = w.Write(data)
}
