package main

import (
	"github.com/tsongpon/listener/route"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting application")

	redPlanetRoute := route.RedPlanetRouter()
	log.Print("The service is ready to listen and serve.")

	log.Fatal(http.ListenAndServe(":5000", redPlanetRoute))
}
