package handeler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/virattGoogle/bookings/pkg/config"
	"github.com/virattGoogle/bookings/pkg/forms"
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
	render.RenderTemplate(w,r, "home.page.tmpl", &models.Templatedata{})

}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,r, "about.page.tmpl", &models.Templatedata{})

}
func Majors(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,r,"major.page.tmpl", &models.Templatedata{})

}

func Generals(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,r,"general.page.tmpl", &models.Templatedata{})

}

func Availability(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,r,"search-availability.page.tmpl", &models.Templatedata{})



}


func PostAvailability(w http.ResponseWriter, r *http.Request){

	start := r.Form.Get("start")
	end:= r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s",start,end)))

}

type jsonResponse struct{
	OK bool `json:"ok"`
	Message string `json:"message"`
}


func AvailabilityJson(w http.ResponseWriter, r *http.Request){

resp := jsonResponse{
	OK:true,
	Message: "Available",
}

out,err := json.MarshalIndent(resp,"","     ")

if err !=nil{
	log.Panicln(err)
}
w.Header().Set("Content-Type", "application/json")
w.Write(out)
}



func Contact(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,r,"about.page.tmpl", &models.Templatedata{})

}

func Reservation(w http.ResponseWriter, r *http.Request){

	render.RenderTemplate(w,r,"make-reservation.page.tmpl", &models.Templatedata{
		Form: forms.New(nil),
	})

}


func PostReservation(w http.ResponseWriter, r *http.Request){

err := r.ParseForm()
if err != nil {
	log.Println(err)
	return

}
reservation := models.Reservation{
	FirstName: r.Form.Get("first_name"),
	LastName: r.Form.Get("last_name"),
	Phone: r.Form.Get("phone"),
	Email: r.Form.Get("email"),
}



form:= forms.New(r.PostForm)

form.Has("first_name",r)

if !form.Valid(){
	data := make(map[string]interface{})
	data["reservation"] = reservation
    render.RenderTemplate(w,r,"make-reservation.page.tmpl", &models.Templatedata{
		Form: form,
		Data: data,
	})

}

}

