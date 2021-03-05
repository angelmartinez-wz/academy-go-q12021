package usecase

import (
	"errors"
	"net/url"
	"strconv"	
	"bootcamp/domain/model"
	"bootcamp/service/csv"
	"bootcamp/service/db"
	"bootcamp/utils"
)

/*
GetPokemonCSV returns a Pokemon list read from a CSV file
If query params exists, filter the Pokemon to return
*/
func GetPokemonCSV(queryParams url.Values) (model.PokemonList, error) {
	pokemonList, err := csv.ReadCSV()

	if err == nil {
		if len(queryParams) > 0 {
			pokemon:= utils.GetPokemonByKey(queryParams, pokemonList)
			var pokemonSubset model.PokemonList
			pokemonSubset = append(pokemonSubset, pokemon)
			return pokemonSubset, nil
		}
	}

	return pokemonList, err
}

/*
GetPokemonCSVById returns a Pokemon read from a CSV file
*/
func GetPokemonCSVById(id string) (model.Pokemon, error) {
	var pokemon model.Pokemon
	pokemonList, err := csv.ReadCSV()

	if pokemonList != nil {
		pokemon = pokemonList[0]
	}

	if err == nil {
		var pokemonSubset model.PokemonList

		if id != "" {
			index, _ := strconv.Atoi(id)
			pokemonId := index - 1
	
			if pokemonId <= len(pokemonList) - 1 {
				pokemonSubset = append(pokemonSubset, pokemonList[pokemonId])
				return pokemonSubset[0], nil
			}

			err = errors.New("Invalid index")
		}
	}

	return pokemon, err
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
func GetPokemonById(id string) (model.Pokemon, error) {
	var pokemon model.Pokemon
	objectId, err := utils.GetObjectIdFromParams(id)

	if err == nil {
		pokemon, err = db.GetPokemonById(objectId)
	}

	return pokemon, err
}

/*
AddPokemon inserts Pokemon information in the database
*/
func AddPokemon(pokemon model.Pokemon) (model.Pokemon, error) {
	pokemon, err := db.AddPokemon(pokemon)
	return pokemon, err
}

/*
UpdatePokemon updates Pokemon information in the database
*/
func UpdatePokemon(id string, pokemon model.Pokemon) (model.Pokemon, error) {
	objectId, err := utils.GetObjectIdFromParams(id)

	if err == nil {
		pokemon, err = db.UpdatePokemon(objectId, pokemon)
	}

	return pokemon, err
}

/*
DeletePokemon deletes Pokemon information in the database
*/
func DeletePokemon(id string) (model.Pokemon, error) {
	var pokemon model.Pokemon
	objectId, err := utils.GetObjectIdFromParams(id)

	if err == nil {
		pokemon, err = db.DeletePokemon(objectId)
	}

	return pokemon, err
}
