package handler

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/acamilleri/slidups/pkg/logger"
)

// Handler - Handler structure
type Handler struct {
	logger            *logger.Logger
	uploadDestination string
}

// New - Initialize handler
func New(logger *logger.Logger, uploadDestination string) *Handler {
	return &Handler{
		logger:            logger,
		uploadDestination: uploadDestination,
	}
}

// RegisteredRoutes - Registered handler routes
func (h *Handler) RegisteredRoutes() {
	http.HandleFunc("/upload", h.UploadHandler)
	http.Handle("/metrics", promhttp.Handler())
}
