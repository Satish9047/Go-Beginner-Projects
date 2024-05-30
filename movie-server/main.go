package main

import (
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

// var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){}
func getMovie(w http.ResponseWriter, r *http.Request){}
func createMovie(w http.ResponseWriter, r *http.Request){}
func updateMovie(w http.ResponseWriter, r *http.Request){}
func deleteMovie(w http.ResponseWriter, r *http.Request){}


func main (){
	r:= mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods(("GET"))
	r.HandleFunc("/movies/{id}", getMovie).Methods(("GET"))
	r.HandleFunc("/movies", createMovie).Methods(("POST"))
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
}

