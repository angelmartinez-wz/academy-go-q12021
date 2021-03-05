package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"reflect"
	"bootcamp/domain/model"
	"gopkg.in/mgo.v2/bson"
)

/*
GetObjectIdFromParams transforms an /{id} to a ObjectId
*/
func GetObjectIdFromParams(id string) (bson.ObjectId, error) {
	var objectId bson.ObjectId

	if id == "" || !bson.IsObjectIdHex(id) {
		return objectId,	errors.New("Invalid id provided")
	}

	objectId = bson.ObjectIdHex(id)
	return objectId, nil
}

/*
GetPokemonFromReader decode Pokemon from reader (body request) and returns the Pokemon information
*/
func GetPokemonFromReader(reader io.ReadCloser) (model.Pokemon, error) {
	var tempPokemon model.Pokemon
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&tempPokemon)

	if err == nil {
		defer reader.Close()		
	}

	return tempPokemon, err
}

func getFieldString(pokemon *model.Pokemon, field string) string {
	r := reflect.ValueOf(pokemon)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

/*
GetPokemonByKey returns a Pokemon filtered by a query param property for the given PokemonList
*/
func GetPokemonByKey(params url.Values, pokemonList model.PokemonList) model.Pokemon {
	var filteredPokemon model.Pokemon
	key := reflect.ValueOf(params).MapKeys()[0].Interface().(string)
	value := params[key][0]
	
	for _, pokemon := range pokemonList {
		if getFieldString(&pokemon, key) == value {
			filteredPokemon = pokemon
			break
		}
	}

	return filteredPokemon
}