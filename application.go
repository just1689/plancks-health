package main

import (
	"fmt"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-brutus/api"
	"git.amabanana.com/plancks-cloud/pc-go-brutus/controller/startup"
	"git.amabanana.com/plancks-cloud/pc-go-brutus/model"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Starting")

	startup.Init()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/health", api.CORSHandler(api.Health)).Methods("GET", "OPTIONS")
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", model.ListentPort), router))
	//This (log.Fatal) is a blocking call, no code beneath here will call
}
