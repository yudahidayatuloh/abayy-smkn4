package handlers

import (
	"html/template"
	"net/http"
	"smkn4-go/models"
)

func EskulHandler(w http.ResponseWriter, r *http.Request) {
	page := struct {
		Title  string
		Eskul  []models.DataEskul
	}{
		Title:  "Ekstrakulikuler",
		Eskul:  models.IsiEskul,
	}

	tmpl, err := template.ParseFiles("templates/eskul.html","templates/layout.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, page); err != nil {
		http.Error(w, "Render Error: "+err.Error(), http.StatusInternalServerError)
	}
}
