package controller

import (
	"fmt"
	"net/http"
	"bootcamp/usecase"
)

/*
HelloWorld prints a hellow world
*/
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msg := usecase.HelloWorld()
	fmt.Fprintf(w, msg)
}