package csv

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"bootcamp/domain/model"
)

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