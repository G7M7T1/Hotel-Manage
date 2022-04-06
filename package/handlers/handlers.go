package handlers

import (
	"hotel-manage/package/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RDTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RDTemplate(w, "about.page.tmpl")
}
