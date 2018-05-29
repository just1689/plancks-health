package main

import (
	"fmt"
	"log"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-brutus/api"
	"github.com/gorilla/mux"
)

const port = 8080

func main() {

	pingFunction()
}

func pingFunction() {
	fmt.Println("Hello World")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/ping", api.CORSHandler(api.Ping)).Methods("GET", "OPTIONS")
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
	return
}
