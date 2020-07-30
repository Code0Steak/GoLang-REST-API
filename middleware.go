package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//Initialize Router
	r := mux.NewRouter()

	//Route Handlers
	r.HandleFunc("/api/recipes", getRecipes).Methods("GET")
	r.HandleFunc("/api/recipes/{id}", getRecipe).Methods("GET")
	r.HandleFunc("/api/recipes", createRecipe).Methods("POST")
	r.HandleFunc("/api/recipes/{id}", updateRecipes).Methods("PUT")
	r.HandleFunc("/api/recipes/{id}", deleteRecipes).Methods("DELETE")

	//Spin-up a server
	log.Fatal(http.ListenAndServe(":5000", r))

}
