package delivery

import (
	"groupie-tracker/models"
	"net/http"
	"text/template"
)

type Handler struct {
	tmpl     *template.Template
	Artists  []models.Artists
	Relation models.Relation
}

func NewHandler() *Handler {
	return &Handler{
		tmpl: template.Must(template.ParseGlob("./ui/templates/*.html")),
	}
}

func (h *Handler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.homePage)
	mux.HandleFunc("/artists/", h.artistPage)
	mux.Handle("/ui/static/", http.StripPrefix("/ui/static/", http.FileServer(http.Dir("./ui/static/"))))
}
