package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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

// get single movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r)
	fmt.Printf(vars["id"])
	id := vars["id"]
	

	for _, movie := range movies {
        if movie.ID == id {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(movie)
            return
        }
    }

}
func createMovie(w http.ResponseWriter, r *http.Request){}
func updateMovie(w http.ResponseWriter, r *http.Request){}
// delete
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies =  append(movies[:index], movies[index+1:]...)
			break
		}
	}
}




func main (){
	r:= mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "123456789", Title: "The Godfather", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "123456790", Title: "The Maze Runner", Director: &Director{Firstname: "Rob", Lastname: "Lucci"}})

	r.HandleFunc("/movies", getMovies).Methods(("GET"))
	r.HandleFunc("/movies/{id}", getMovie).Methods(("GET"))
	r.HandleFunc("/movies", createMovie).Methods(("POST"))
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server is running on port 8080\n")
	if err:= http.ListenAndServe(":8080", r); err!=nil {
		log.Fatal(err)
	}
}

