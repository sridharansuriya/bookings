package handlers

import (
	"net/http"

	"github.com/sridharansuriya/bookings/pkg/config"
	"github.com/sridharansuriya/bookings/pkg/models"
	"github.com/sridharansuriya/bookings/pkg/render"
)

var Repo *Repository

func SetAppConfig(app *config.AppConfig) *Repository {
	return &Repository{
		AppConfig: app,
	}
}

func SetRepo(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (repo *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	repo.AppConfig.Session.Put(r.Context(), "remote_ip", r.RemoteAddr)
	render.RenderTemplate(rw, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (repo *Repository) About(rw http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	remoteIP := repo.AppConfig.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(rw, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

type Repository struct {
	AppConfig *config.AppConfig
}
