package main

import (
	"github.com/tsongpon/listener/route"
	"log"
	"net/http"
	"github.com/tsongpon/listener/data"
)

func main() {
	log.Print("Starting application")
	data.InitDB()
	defer data.CloseDB()
	redPlanetRoute := route.NewRedPlanetRouter()
	log.Print("The service is ready to listen and serve.")

	log.Fatal(http.ListenAndServe(":5000", redPlanetRoute))
}
