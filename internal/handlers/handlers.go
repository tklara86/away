package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/tklara86/away/internal/config"
	"github.com/tklara86/away/internal/driver"
	"github.com/tklara86/away/internal/forms"
	"github.com/tklara86/away/internal/helpers"
	"github.com/tklara86/away/internal/models"
	"github.com/tklara86/away/internal/render"
	"github.com/tklara86/away/internal/repository"
	"github.com/tklara86/away/internal/repository/dbrepo"
	"net/http"
	"strconv"
	"time"
)


// Repo the repository uses by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB: dbrepo.NewPostgresRepo(db.SQL, a),
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

	render.Template(w, r,"home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "About page"

	render.Template(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// RoseHill is the rosehill page handler
func (m *Repository) RoseHill(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "RoseHill"

	render.Template(w, r, "rosehill.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}


// MillHouse is the millhouse page handler
func (m *Repository) MillHouse(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "MillHouse"

	render.Template(w, r, "millhouse.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Availability is the availability page handler
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "Search"

	render.Template(w, r,"search-availability.page.tmpl", &models.TemplateData{
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
		helpers.ServerError(w, err)
		return
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

	render.Template(w, r, "contact.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
// MakeReservation is the reservation page handler
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "Make a Reservation"

	var emptyReservation models.Reservation

	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		Form: forms.New(nil),
		Data: data,
	})
}

// PostMakeReservation handles the posting of a reservation form
func (m *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")

	// 2020-01-01 01/02 03:04:05PM '06 -0700

	layout := "2006-01-02"
	startDate, err := time.Parse(layout,sd)
	if err != nil {
		helpers.ServerError(w, err)
	}
	endDate, err := time.Parse(layout,ed)
	if err != nil {
		helpers.ServerError(w, err)
	}

	roomId, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(w, err)
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("firstName"),
		LastName: r.Form.Get("lastName"),
		Email: r.Form.Get("email"),
		Phone: r.Form.Get("phone"),
		StartDate: startDate,
		EndDate: endDate,
		RoomID: roomId,
	}

	form := forms.New(r.PostForm)

	form.Required("firstName", "lastName", "email", "phone")
	form.MinLength("firstName", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	err = m.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(w, err)
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}


// ReservationSummary shows reservation summary
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pageTitle"] = "Reservation Summary"

	reservation, ok := m.App.Session.Get(r.Context(),"reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("can't get error from session")
		m.App.Session.Put(r.Context(), "error", "Cannot get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.Template(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		Data: data,
	})



}