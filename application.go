package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/plancks-cloud/plancks-health/api"
	"github.com/plancks-cloud/plancks-health/controller/startup"
	"github.com/plancks-cloud/plancks-health/model"
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
