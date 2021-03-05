package controller

import (
	"net/http"
	"bootcamp/usecase"
	"bootcamp/service/network"
	"github.com/gorilla/mux"
)

/*
GetPokemonCSV returns a JSON with a Pokemon  list information
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
GetPokemonCSVById returns a JSON with the Pokemon information
*/
func GetPokemonCSVById(w http.ResponseWriter, r *http.Request) {
	pokemon, err := usecase.GetPokemonCSVById(r)
	network.Response(w, pokemon, err)
}

/*
GetPokemon returns a JSON Pokemon list
*/
func GetPokemon(w http.ResponseWriter, r *http.Request) {
	pokemonList, err := usecase.GetPokemon()
	network.ResponseList(w, pokemonList, err)
}

/*
GetPokemonById returns a JSON with the Pokemon information
*/
func GetPokemonById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pokemon, err := usecase.GetPokemonById(params)
	network.Response(w, pokemon, err)
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