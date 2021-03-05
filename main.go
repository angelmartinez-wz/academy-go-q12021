package main

import (
	"fmt"
	"log"
	"net/http"
	"bootcamp/config"
	"bootcamp/router"
)

func main() {
	rt := router.NewRouter()
	configuration := config.GetConfiguration()

	fmt.Println("Server listening in port" + configuration.Port)

	server := http.ListenAndServe(configuration.Port, rt)

	log.Fatal(server)
}
