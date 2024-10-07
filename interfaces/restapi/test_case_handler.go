package restapi

import (
	"encoding/json"
	"net/http"
	"rocket-refactor-with-DDD/application"
)

type RocketHandler struct {
	appService *application.RocketAppService
}

func NewRocketHandler(appService *application.RocketAppService) *RocketHandler {
	return &RocketHandler{
		appService: appService,
	}
}

func (h *RocketHandler) UpdateRocketCapacity(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID       string `json:"id"`
		Capacity int    `json:"capacity"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.appService.UpdateRocketCapacity(req.ID, req.Capacity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
