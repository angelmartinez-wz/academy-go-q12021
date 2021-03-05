package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bootcamp/domain/model"
	"bootcamp/usecase"
	"bootcamp/utils"
	"github.com/gorilla/mux"
)

func setHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return w
}

func validateError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, err.Error())
}

func response(w http.ResponseWriter, pokemon model.Pokemon, err error) {
	if err != nil {
		validateError(w, err)
		return
	}

	w = setHeaders(w)
	json.NewEncoder(w).Encode(pokemon)	
}

func responseList(w http.ResponseWriter, pokemonList model.PokemonList, err error) {
	if err != nil {
		validateError(w, err)
		return
	}

	w = setHeaders(w)
	json.NewEncoder(w).Encode(pokemonList)
}

/*
GetPokemonCSV returns a JSON with a Pokemon  list information
If URL contains a query params look for a Pokemon that matches with that search filter
*/
func GetPokemonCSV(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	pokemonList, err := usecase.GetPokemonCSV(queryParams)

	if len(pokemonList) == 1 {
		response(w, pokemonList[0], err)
		return
	}

	responseList(w, pokemonList, err)
}

/*
GetPokemonCSVById returns a JSON with the Pokemon information
*/
func GetPokemonCSVById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id:= params["id"]
	pokemon, err := usecase.GetPokemonCSVById(id)
	response(w, pokemon, err)
}

/*
GetPokemon returns a JSON Pokemon list
*/
func GetPokemon(w http.ResponseWriter, r *http.Request) {
	pokemonList, err := usecase.GetPokemon()
	responseList(w, pokemonList, err)
}

/*
GetPokemonById returns a JSON with the Pokemon information
*/
func GetPokemonById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	pokemon, err := usecase.GetPokemonById(id)
	response(w, pokemon, err)
}

/*
AddPokemon returns a JSON Pokemon struct with the new added Pokemon information
*/
func AddPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := utils.GetPokemonFromReader(r.Body)
	
	if err == nil {
		pokemon, err = usecase.AddPokemon(pokemon)
	}

	response(w, pokemon, err)
}

/*
UpdatePokemon returns a JSON Pokemon struct with the updated Pokemon information
*/
func UpdatePokemon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	pokemon, err := utils.GetPokemonFromReader(r.Body)
	
	if err == nil {
		pokemon, err = usecase.UpdatePokemon(id, pokemon)
	}

	response(w, pokemon, err)
}

/*
DeletePokemon returns a JSON Pokemon struct with the deleted Pokemon information
*/
func DeletePokemon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	pokemon, err := usecase.DeletePokemon(id)
	response(w, pokemon, err)
}