package main

import (
	"fmt"
	"log"
	"net/http"
	"smkn4-go/handlers"
)

func main() {
	http.HandleFunc("/", handlers.BerandaHandler)
	http.HandleFunc("/profil", handlers.ProfilHandler)
	http.HandleFunc("/eskul", handlers.EskulHandler)
	http.HandleFunc("/galeri", handlers.GaleriHandler)

	fmt.Println("Server siap Dunungan: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
