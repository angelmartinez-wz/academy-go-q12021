package controller

import (
	"net/http"
	"bootcamp/usecase"
	"bootcamp/service/network"
	"github.com/gorilla/mux"
)

/*
GetPokemonCSV returns a JSON with the Pokemon information
If URL not contains /{id} nor query params return a Pokemon array
If URL contains /{id} return the Pokemon for the given index
If URL contains a query params look for a Pokemon that matches with that search filter
*/
func GetPokemonCSV(w http.ResponseWriter, r *http.Request) {
	pokemonList, err := usecase.GetPokemonCSV(r)

	if len(pokemonList) == 1 {
		network.Response(w, pokemonList[0], err)
		return
	}

	network.ResponseList(w, pokemonList, err)
}

/*
GetPokemon returns a JSON Pokemon array or a Pokemon information
If URL not contains /{id} returns a Pokemon array
If URL contains /{id} return the Pokemon for the given index
*/
func GetPokemon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if id := params["id"]; id != "" {
		pokemon, err := usecase.GetPokemonById(params)
		network.Response(w, pokemon, err)
	} else {
		pokemonList, err := usecase.GetPokemon()
		network.ResponseList(w, pokemonList, err)
	}
}

/*
AddPokemon returns a JSON Pokemon struct with the new added Pokemon information
*/
func AddPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := usecase.AddPokemon(r.Body)
	network.Response(w, pokemon, err)
}

/*
UpdatePokemon returns a JSON Pokemon struct with the updated Pokemon information
*/
func UpdatePokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := usecase.UpdatePokemon(mux.Vars(r), r.Body)
	network.Response(w, pokemon, err)
}

/*
DeletePokemon returns a JSON Pokemon struct with the deleted Pokemon information
*/
func DeletePokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := usecase.DeletePokemon(mux.Vars(r))
	network.Response(w, pokemon, err)
}