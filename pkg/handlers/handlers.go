package handlers

import (
	"fmt"
	"net/http"

	"github.com/Doaa-Ismail/go_course/booking/pkg/config"
	"github.com/Doaa-Ismail/go_course/booking/pkg/functions"
	"github.com/Doaa-Ismail/go_course/booking/pkg/models"
	"github.com/Doaa-Ismail/go_course/booking/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplates(w, "home.page.html", &models.TemplateData{})

	//fmt.Printf("r", r)
	//_, _ = fmt.Fprintf(w, "Hello From Home Page ^^ ")
	// n, err := fmt.Fprintf(w, "Hello world")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("Number of Bytes: %d\n", n)
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplates(w, "about.page.html", &models.TemplateData{})
	_, _ = fmt.Fprintf(w, "Hello From About Page ^^ \n")
	//sumfunc(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintln("The Sum Function is called ", functions.Sumfunc(2, 2)))
	//var s string
	//fmt.Sprintf("%s", &s)
	//fmt.Println("s : ", s)
}

func (m *Repository) Divide(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "check.page.html", &models.TemplateData{})
	_, _ = fmt.Fprintf(w, "Hello From Check Page ^^ \n")

	var x, y float32 = 100.0, 0.500
	f, err := functions.DivideNum(x, y)
	if err != nil {
		fmt.Fprintf(w, "Error (%f / %f): Cannot divide by zero\n", x, y)
	}
	fmt.Fprintf(w, "The result of %f / %f = %f\n", x, y, f)

}

func (m *Repository) Hello(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again from Template"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplates(w, "hello.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
