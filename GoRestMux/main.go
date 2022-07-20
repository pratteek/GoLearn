package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{

	{ID: "A1", Title: "Summer", Artist: "Harry", Price: 79.99},
	{ID: "A2", Title: "Winter", Artist: "Zayn", Price: 199.99},
	{ID: "A3", Title: "Spring", Artist: "Niall", Price: 49.99},
	{ID: "A4", Title: "Monsoon", Artist: "Liam", Price: 39.99},
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)
}
func getAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for _, album := range albums {
		if album.ID == id {
			json.NewEncoder(w).Encode(album)
		}
	}
}
func createAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newAlbum album
	if err := json.NewDecoder(r.Body).Decode(&newAlbum); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoding response object", http.StatusBadRequest)
		return
	}
	albums = append(albums, newAlbum)

	response, err := json.Marshal(&albums)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoding response object", http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page")

}

func main() {
	r := mux.NewRouter()
	r.Path("/").Methods(http.MethodGet).HandlerFunc(Home)
	r.Path("/albums").Methods(http.MethodGet).HandlerFunc(getAlbums)
	r.Path("/albums/{id}").Methods(http.MethodGet).HandlerFunc(getAlbum)
	r.Path("/albums").Methods(http.MethodPost).HandlerFunc(createAlbum)

	log.Fatal(http.ListenAndServe(":8080", r))

}
