package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Recipes Struct (Model/Resource)

type Recipes struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Ingredients map[string]string `json:"ingredients"`
	Procedure   map[int]string    `json:"procedure"`
	Cook        string            `json:"recipe_by"`
	Rating      float32           `json:"rating"`
	recipeType  string            `json:"cuisine"`
}

var recipes []Recipes

/* Funcs. for the "GET" method */
func getRecipes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(recipes)

}

func getRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, recipe := range recipes {
		if recipe.ID == params["id"] {
			json.NewEncoder(w).Encode(recipe)
		}
	}

	json.NewEncoder(w).Encode(&Recipes{})

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

	//Fake DB

	//Entry #1
	ingr := map[string]string{
		"🥚":     "four",
		"🍅":     "three and a 1/2",
		"🧅":     "two big",
		"clove": "two",
		"🧀":     "optional",
	}
	procedure := map[int]string{
		1: "2 cloves pressed over a knife. Diced onion in a pan containing heated oil",
		2: "Turmeric and 🌶 to it",
		3: "Crushed 🍅's and pepper in. For 20 mins, steam",
		4: "Make wells to embed 🥚. Cook for 5",
		5: "Serve Hot ♨!",
	}
	recipes = append(recipes, Recipes{

		ID:          "1",
		Name:        "Shakshuka",
		Ingredients: ingr,
		Procedure:   procedure,
		Cook:        "Binging wt. Babish",
		Rating:      4.8,
		recipeType:  "Indian",
	})

	//#2
	ingr2 := map[string]string{
		"Maggie Noodle Packet": "As many",
	}

	procedure2 := map[int]string{
		1: "Follow the instructions on the Maggie Packet 😄",
		2: "Serve Hot and with topping of some veggies(and 🧀)",
	}

	recipes = append(recipes, Recipes{

		ID:          "2",
		Name:        "Maggie Noodle",
		Ingredients: ingr2,
		Procedure:   procedure2,
		Cook:        "You",
		Rating:      5.0,
		recipeType:  "Indian",
	})

	//Route Handlers
	r.HandleFunc("/api/recipes", getRecipes).Methods("GET")
	r.HandleFunc("/api/recipes/{id}", getRecipe).Methods("GET")
	r.HandleFunc("/api/recipes", createRecipe).Methods("POST")
	r.HandleFunc("/api/recipes/{id}", updateRecipe).Methods("PUT")
	r.HandleFunc("/api/recipes/{id}", deleteRecipe).Methods("DELETE")

	//Spin-up a server
	log.Fatal(http.ListenAndServe(":5000", r))

}
