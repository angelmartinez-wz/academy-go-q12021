package router

import (
	"net/http"
	"bootcamp/controller"
	"bootcamp/domain/model"
	"github.com/gorilla/mux"
)

/*
NewRouter implements a gorilla/mux router with the routes of the API
*/
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}
	return router
}

var pokemonPath = "/pokemon"
var id = "/{id}"
var pokemonPathWithId = pokemonPath + id
var csvPokemonPath = "/csv" + pokemonPath

var routes = model.Routes{
	model.Route{
		"HelloWorld",
		http.MethodGet,
		"/",
		controller.HelloWorld,
	},
	model.Route{
		"GetPokemonListCsv",
		http.MethodGet,
		csvPokemonPath,
		controller.GetPokemonCSV,
	},
	model.Route{
		"GetPokemonCsv",
		http.MethodGet,
		csvPokemonPath + id,
		controller.GetPokemonCSVById,
	},
	model.Route{
		"AddPokemon",
		http.MethodPost,
		pokemonPath,
		controller.AddPokemon,
	},
	model.Route{
		"GetPokemon",
		http.MethodGet,
		pokemonPath,
		controller.GetPokemon,
	},
	model.Route{
		"GetPokemonById",
		http.MethodGet,
		pokemonPathWithId,
		controller.GetPokemonById,
	},
	model.Route{
		"UpdatePokemon",
		http.MethodPut,
		pokemonPathWithId,
		controller.UpdatePokemon,
	},
	model.Route{
		"DeletePokemon",
		http.MethodDelete,
		pokemonPathWithId,
		controller.DeletePokemon,
	},
}