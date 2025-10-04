package handlers

import (
	"html/template"
	"net/http"
	"smkn4-go/models"
)

func GaleriHandler(w http.ResponseWriter, r *http.Request) {
	page := struct {
		Title  string
		Galeri []models.DataGaleri
	}{
		Title:  "Galeri SMKN 4 Tasikmalaya",
		Galeri: models.IsiGaleri,
	}

	tmpl, err := template.ParseFiles("templates/galeri.html","templates/layout.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, page); err != nil {
		http.Error(w, "Render error: "+err.Error(), http.StatusInternalServerError)
	}
}
