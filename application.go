package main

import (
	"fmt"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-brutus/api"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const port = 8080

func main() {

	log.Info("Starting")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/health", api.CORSHandler(api.Health)).Methods("GET", "OPTIONS")
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), router))
}
