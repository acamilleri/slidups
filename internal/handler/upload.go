package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/acamilleri/slidups/internal/model"
)

// UploadHandler - Handler to handle slide upload request
func (h *Handler) UploadHandler(w http.ResponseWriter, req *http.Request) {
	r := model.Request{Request: req}
	defer r.Body.Close()

	file, err := r.GetFile()
	if err != nil {
		h.NewResponse(w, http.StatusBadRequest, "invalid payload")
		h.logger.WithError(err).Error("failed to upload file")
		return
	}

	fileDestination := h.uploadDestination
	if file.Destination != "" {
		fileDestination = file.Destination
	}

	err = os.MkdirAll(fileDestination, os.FileMode(0775))
	if err != nil {
		h.NewResponse(w, http.StatusBadRequest, "invalid directory path")
		h.logger.WithError(err).Error("invalid directory path")
		return
	}

	dest := fmt.Sprintf("%s/%s", fileDestination, file.Name)
	err = ioutil.WriteFile(dest, file.Content, os.FileMode(0775))
	if err != nil {
		h.NewResponse(w, http.StatusInternalServerError, "failed to write file")
		h.logger.WithError(err).Error("failed to write file")
		return
	}

	h.logger.Infof("file %s received and write to %s", file.Name, dest)
	h.NewResponse(w, http.StatusOK, "ok")
}
