package main

import (
	"fmt"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-brutus/api"
	"git.amabanana.com/plancks-cloud/pc-go-brutus/controller/startup"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const port = 8080

func main() {

	log.Info("Starting")

	startup.Init()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/health", api.CORSHandler(api.Health)).Methods("GET", "OPTIONS")
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
	//This (log.Fatal) is a blocking call, no code beneath here will call
}
