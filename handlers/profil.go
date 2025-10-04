package handlers

import (
	"html/template"
	"net/http"
	"smkn4-go/models"
)

func ProfilHandler(w http.ResponseWriter, r *http.Request) {
	page := struct {
		Title  string
		Profil models.DataProfil
	}{
		Title:  "Profil Sekolah",
		Profil: models.IsiProfil,
	}
	tmpl, err := template.ParseFiles("templates/profil.html","templates/layout.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, page); err != nil {
		http.Error(w, "Render error: "+err.Error(), http.StatusInternalServerError)
	}
}
