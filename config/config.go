package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Port  			string
	MongoUrl  	string
	DbName			string
	Collection	string
}

func GetConfiguration() *Configuration {
	file, _ := os.Open(".env")

	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := &Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		log.Fatal("Could not initialize envars")
	}

	return configuration
}