package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"strconv"
// 	"github.com/gorilla/mux"
// )


type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

// get all movies
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request){}
func createMovie(w http.ResponseWriter, r *http.Request){}
func updateMovie(w http.ResponseWriter, r *http.Request){}
func deleteMovie(w http.ResponseWriter, r *http.Request){}


func main (){
	r:= mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "123456789", Title: "The Godfather", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "123456790", Title: "The Maze Runner", Director: &Director{Firstname: "Rob", Lastname: "Lucci"}})

	r.HandleFunc("/movies", getMovies).Methods(("GET"))
	r.HandleFunc("/movies/{id}", getMovie).Methods(("GET"))
	r.HandleFunc("/movies", createMovie).Methods(("POST"))
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server is running on port 8080")
	if err:= http.ListenAndServe(":8080", r); err!=nil {
		log.Fatal(err)
	}
}

