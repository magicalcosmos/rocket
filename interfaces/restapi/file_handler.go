package restapi

import (
	"encoding/json"
	"net/http"
	"project-with-ddd/application"
)

type FileHandler struct {
	fileAppService *application.FileAppService
}

func NewFileHandler(appService *application.FileAppService) *FileHandler {
	return &FileHandler{
		fileAppService: appService,
	}
}

func (h *FileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Content []byte `json:"content"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	err := h.fileAppService.UploadFile(req.ID, req.Name, req.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
