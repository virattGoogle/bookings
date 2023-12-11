package handeler

import (
	"net/http"
	"github.com/virattGoogle/bookings/pkg/config"
	"github.com/virattGoogle/bookings/pkg/models"
	"github.com/virattGoogle/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.Appconfig //Appconfig is a struct that has map[string]*template.Template
}

func NewRepo(a *config.Appconfig) *Repository { //converts config.Appconfig to repository
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) { //Sets Repo for new Handlers
	Repo = r
}

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.Templatedata{})

}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.Templatedata{})

}
func Majors(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,"major.page.tmpl", &models.Templatedata{})

}

func Generals(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,"general.page.tmpl", &models.Templatedata{})

}

func Availability(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,"search-availability.page.tmpl", &models.Templatedata{})

}
func Contact(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,"about.page.tmpl", &models.Templatedata{})

}

func Reservation(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,"make-reservation.page.tmpl", &models.Templatedata{})

}