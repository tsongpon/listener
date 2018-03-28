package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Starting application")

	redPlanetRoute := RedPlanetRouter()
	log.Print("The service is ready to listen and serve.")

	log.Fatal(http.ListenAndServe(":5000", redPlanetRoute))
}
