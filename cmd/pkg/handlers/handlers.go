package handlers

import (
	"github.com/tklara86/away/cmd/pkg/config"
	"github.com/tklara86/away/cmd/pkg/models"
	"github.com/tklara86/away/cmd/pkg/render"
	"net/http"
)


// Repo the repository uses by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "Home page"

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "About page"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// RoseHill is the rosehill page handler
func (m *Repository) RoseHill(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "RoseHill"

	render.RenderTemplate(w, "rosehill.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}


// MillHouse is the millhouse page handler
func (m *Repository) MillHouse(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "MillHouse"

	render.RenderTemplate(w, "millhouse.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Search is the search availability page handler
func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "Search"

	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Contact is the contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "Contact"

	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
