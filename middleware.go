package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Recipes Struct (Model/Resource)

type Recipes struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Ingredients map[string]string `json:"ingredients"`
	Procedure   map[string]string `json:"procedure"`
	Cook        string            `json:"recipe_by"`
	Rating      string            `json:"rating"`
	recipeType  string            `json:"cuisine"`
}

var recipes []Recipes

/* Funcs. for the "GET" method */
func getRecipes(w http.ResponseWriter, r *http.Request) {

}

func getRecipe(w http.ResponseWriter, r *http.Request) {

}

/* Funcs. for the "POST" method */
func createRecipe(w http.ResponseWriter, r *http.Request) {

}

/* Funcs. for the "PUT" method */
func updateRecipe(w http.ResponseWriter, r *http.Request) {

}

/* Funcs. for the "DELETE" method */
func deleteRecipe(w http.ResponseWriter, r *http.Request) {

}

func main() {

	//Initialize Router
	r := mux.NewRouter()

	//Route Handlers
	r.HandleFunc("/api/recipes", getRecipes).Methods("GET")
	r.HandleFunc("/api/recipes/{id}", getRecipe).Methods("GET")
	r.HandleFunc("/api/recipes", createRecipe).Methods("POST")
	r.HandleFunc("/api/recipes/{id}", updateRecipe).Methods("PUT")
	r.HandleFunc("/api/recipes/{id}", deleteRecipe).Methods("DELETE")

	//Spin-up a server
	log.Fatal(http.ListenAndServe(":5000", r))

}
