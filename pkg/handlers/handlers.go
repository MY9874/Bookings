package handlers

import (
	"net/http"

	"github.com/moyu/bookings/pkg/config"
	"github.com/moyu/bookings/pkg/models"
	"github.com/moyu/bookings/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository, starts only with the app config
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler, has a receiver that have access to everything inside repository (application config)
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler, has a receiver that have access to everything inside repository (application config))
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic, say get some data
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	//send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
