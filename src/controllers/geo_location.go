package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Locally-build/acme-weather-backend/services"
	"github.com/gorilla/mux"
)

func GeoLocationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	city := params["city"]
	state := params["state"]
	country := params["country"]

	if state != "" {
		log.Printf("[GET] Geographic coordinates for the %s city in the state of %s on %s", city, state, country)
	} else {
		log.Printf("[GET] Geographic coordinates for the %s city on %s", city, country)
	}

	svc := services.GetWeatherService()
	resp, err := svc.GetCityCoordinates(city, state, country)
	if err != nil {
		log.Print(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Print(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Print(err)
	}
}
