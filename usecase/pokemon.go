package usecase

import (
	"errors"
	"io"
	"net/http"
	"strconv"	
	"bootcamp/domain/model"
	"bootcamp/service/db"
	"bootcamp/utils"
	"github.com/gorilla/mux"
)

/*
GetPokemonCSV returns a Pokemon list read from a CSV file
If query params exists or /{id}, filter the Pokemon
Otherwise, returns the list
*/
func GetPokemonCSV(r *http.Request) (model.PokemonList, error) {
	pokemonList, err := utils.ReadCSV()
	var pokemonSubset model.PokemonList
	params := mux.Vars(r)

	if err == nil {
		id:= params["id"]

		if id != "" {
			index, _ := strconv.Atoi(id)
			pokemonId := index - 1
	
			if pokemonId <= len(pokemonList) - 1 {
				pokemonSubset = append(pokemonSubset, pokemonList[pokemonId])
				return pokemonSubset, nil
			}

			err = errors.New("Invalid index")
			return nil, err
		}
	}

	queryParams := r.URL.Query()

	if len(queryParams) > 0 {
		pokemon:= utils.GetPokemonByKey(queryParams, pokemonList)
		var pokemonSubset model.PokemonList
		pokemonSubset = append(pokemonSubset, pokemon)
		return pokemonSubset, nil
	}

	return pokemonList, err
}

/*
GetPokemon retrieve all existent Pokemon from the database
*/
func GetPokemon() (model.PokemonList, error) {
	pokemonList, err := db.GetPokemon()
	return pokemonList, err
}

/*
GetPokemonById retrieves Pokemon information that matches with a given id from the database
*/
func GetPokemonById(params map[string]string) (model.Pokemon, error) {
	var pokemon model.Pokemon
	objectId, err := utils.GetObjectIdFromParams(params["id"])

	if err == nil {
		pokemon, err = db.GetPokemonById(objectId)
	}

	return pokemon, err
}

/*
AddPokemon inserts Pokemon information in the database
*/
func AddPokemon(reader io.ReadCloser) (model.Pokemon, error) {
	pokemon, err := utils.GetPokemonFromReader(reader)
	
	if err == nil {
		pokemon, err = db.AddPokemon(pokemon)
	}

	return pokemon, err
}

/*
UpdatePokemon updates Pokemon information in the database
*/
func UpdatePokemon(params map[string]string, reader io.ReadCloser) (model.Pokemon, error) {
	var pokemon model.Pokemon
	pokemon, err := utils.GetPokemonFromReader(reader)
	objectId, err := utils.GetObjectIdFromParams(params["id"])

	if err == nil {
		pokemon, err = db.UpdatePokemon(objectId, pokemon)
	}

	return pokemon, err
}

/*
DeletePokemon deletes Pokemon information in the database
*/
func DeletePokemon(params map[string]string) (model.Pokemon, error) {
	var pokemon model.Pokemon
	objectId, err := utils.GetObjectIdFromParams(params["id"])

	if err == nil {
		pokemon, err = db.DeletePokemon(objectId)
	}

	return pokemon, err
}
