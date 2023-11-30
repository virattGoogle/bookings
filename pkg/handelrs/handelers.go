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



    
	stringMap := make(map[string]string)

	stringMap["test"] = "hello Vinayak Rattan from Handler"
	render.RenderTemplate(w, "home.page.tmpl", &models.Templatedata{
		StringMap: stringMap,
	})

}

func About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello Vinayak Rattan from Handler About"
	render.RenderTemplate(w, "about.page.tmpl", &models.Templatedata{
		StringMap: stringMap,
	})

}
