package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola desde func home"))
}
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/snippet/view/{id}", snippetView)
	http.HandleFunc("/snippet/create", snippetCreate)

	log.Print("El server corre on :4000")

	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}