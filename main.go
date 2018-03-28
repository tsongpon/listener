package main

import (
	"log"
	"net/http"
	"github.com/tsongpon/listener/logger"
)

func main() {
	log.Print("Starting application")
	logger.Loged("testing")

	redPlanetRoute := RedPlanetRouter()
	log.Print("The service is ready to listen and serve.")

	log.Fatal(http.ListenAndServe(":5000", redPlanetRoute))
}
