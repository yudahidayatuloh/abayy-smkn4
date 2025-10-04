package handlers

import (
	"html/template"
	"net/http"
	"smkn4-go/models"
	"smkn4-go/utils"
)



func BerandaHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	var jurusanAmbilWarna []models.DataJurusan
	for _, j := range models.IsiJurusan {
		j.Warna = utils.GetWarna(j.NamaJ)
		jurusanAmbilWarna = append(jurusanAmbilWarna, j)
	}

	page := struct {
		Title         string
		Jurusan       []models.DataJurusan
		JumlahJurusan int
		JumlahSiswa int
		JumlahEskul int
	}{
		Title:         "Beranda",
		Jurusan:       jurusanAmbilWarna,
		JumlahJurusan: len(jurusanAmbilWarna),
		JumlahSiswa: len(models.IsiSiswa),
		JumlahEskul: len(models.IsiEskul),
	}

	tmpl, err := template.ParseFiles("templates/beranda.html","templates/layout.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, page); err != nil {
		http.Error(w, "Render Error: "+err.Error(), http.StatusInternalServerError)
	}
}
