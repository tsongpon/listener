package main

import (
	"github.com/tsongpon/listener/route"
	"log"
	"net/http"
	"github.com/tsongpon/listener/data"
	"github.com/tsongpon/listener/config"
)

func main() {
	log.Print("Starting application")
	data.InitDB(config.GetDBHost(), config.GetDBName())
	defer data.CloseDB()
	redPlanetRoute := route.NewRedPlanetRouter()
	log.Print("The service is ready to listen and serve.")

	log.Fatal(http.ListenAndServe(":5000", redPlanetRoute))
}
