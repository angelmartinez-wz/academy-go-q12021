package utils

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"
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

/*
ReadCSV read a CSV with Pokemon information and transform the content to the Pokemon struct type
*/
func ReadCSV() (model.PokemonList, error) {
	var pokemonList model.PokemonList
	recordFile, err := os.Open("assets/pokemon.csv")

	if err != nil {
		return nil, errors.New("Could not open file")
	}

	reader := csv.NewReader(recordFile)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, errors.New("Could not read file")
	}

	for _, pokemon := range records {
		id, err := strconv.Atoi(pokemon[0])
		
		if err != nil {
			return pokemonList, errors.New("Cannot get id from row")
		}
				
		pkNumber, err := strconv.Atoi(pokemon[1])
		
		if err != nil {
			return pokemonList, errors.New("Cannot get pokedex number from row")
		}
		
		pk := model.Pokemon{Id:id, PokedexNumber: pkNumber, Name:pokemon[2], Types:pokemon[3], Region:pokemon[4]}
		pokemonList = append(pokemonList, pk)
	}
		
	err = recordFile.Close()
		
	if err != nil {
		return pokemonList, errors.New("Error while closing file")
	}
		
	return pokemonList, nil
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