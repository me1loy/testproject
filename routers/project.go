package routers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.Repo3.GetProjects(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch projects", http.StatusInternalServerError)
		return
	}

	if projects == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("[]"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
