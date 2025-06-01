package routers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"server/dto"
	"strconv"
)

type SupportPageData struct {
	Error   string
	Success string
}

func RenderSupportPage(w http.ResponseWriter, data SupportPageData) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Printf("template.ParseFiles error: %v", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("ParseForm: %v", err)
		RenderSupportPage(w, SupportPageData{Error: "Ungültige Formulardaten"})
		return
	}

	projectID, err := strconv.Atoi(r.FormValue("project_id"))
	if err != nil || projectID == 0 {
		RenderSupportPage(w, SupportPageData{Error: "Bitte wähle ein gültiges Projekt"})
		return
	}

	dtoReq := dto.CreateUserRequest{
		Name:      r.FormValue("name"),
		Email:     r.FormValue("email"),
		ProjectID: projectID,
		Help:      r.FormValue("help"),
		Message:   r.FormValue("message"),
	}

	if dtoReq.Name == "" || dtoReq.Email == "" {
		RenderSupportPage(w, SupportPageData{Error: "Name und E-Mail sind erforderlich"})
		return
	}

	userModel := dtoReq.ToModel()

	created, err := h.Repo3.CreateUser(r.Context(), userModel)
	if err != nil {
		log.Printf("DB error: %v", err)
		RenderSupportPage(w, SupportPageData{Error: "Fehler beim Speichern. Bitte versuche es später erneut"})
		return
	}

	RenderSupportPage(w, SupportPageData{
		Success: fmt.Sprintf("Danke, %s! Wir haben deine Nachricht erhalten.", created.Name),
	})
}
