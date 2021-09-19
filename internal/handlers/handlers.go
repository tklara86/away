package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/tklara86/away/internal/config"
	"github.com/tklara86/away/internal/models"
	"github.com/tklara86/away/internal/render"
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

	render.RenderTemplate(w, r,"home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "About page"

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// RoseHill is the rosehill page handler
func (m *Repository) RoseHill(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "RoseHill"

	render.RenderTemplate(w, r, "rosehill.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}


// MillHouse is the millhouse page handler
func (m *Repository) MillHouse(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "MillHouse"

	render.RenderTemplate(w, r, "millhouse.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Availability is the availability page handler
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "Search"

	render.RenderTemplate(w, r,"search-availability.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// PostAvailability is the search availability page handler
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	startDate := r.Form.Get("startDate")
	endDate := r.Form.Get("endDate")

	_, err := w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", startDate,
		endDate)))
	if err != nil {
		fmt.Println(err)
	}
}


type jsonResponse struct {
	OK bool 			`json:"ok"`
	Message string 		`json:"message"`
}
// AvailabilityJSON handles request for availability and sends JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	response := jsonResponse{
		OK: true,
		Message: "Available",
	}

	j, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(j)
	if err != nil {
		fmt.Println(err.Error())
	}
}


// Contact is the contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "Contact"

	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
// MakeReservation is the reservation page handler
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "Make a Reservation"

	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
