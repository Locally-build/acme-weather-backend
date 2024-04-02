package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Parallels/acme-weather-demo-backend/entities"
	"github.com/Parallels/acme-weather-demo-backend/services"
	"github.com/gorilla/mux"
)

func CityWeatherForecastHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	city := params["city"]
	state := params["state"]
	country := params["country"]
	osSystem := r.URL.Query().Get("os")

	log.Printf("osSystem: %s", osSystem)

	log.Printf("Getting weather for %s", city)
	svc := services.GetWeatherService()
	resp, err := svc.GetCityWeatherForecast(city, state, country)
	if err != nil {
		log.Print(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	forecastResponse := entities.ForecastApiResponse{
		City: resp.City.Name,
	}

	forecastResponse.List = make([]entities.Weather, 0)
	for _, item := range resp.List {
		weatherItem := entities.Weather{}
		weatherItem.DateTime = item.DtTxt
		weatherItem.Temperature = fmt.Sprintf("%.0f", item.Main.Temp)
		weatherItem.Description = item.Weather[0].Description
		weatherItem.FeelsLike = fmt.Sprintf("%.0f", item.Main.FeelsLike)
		weatherItem.Humidity = fmt.Sprintf("%d", item.Main.Humidity)
		weatherItem.WindSpeed = fmt.Sprintf("%.0f", item.Wind.Speed)
		weatherItem.Pressure = fmt.Sprintf("%d", item.Main.Pressure)
		weatherItem.MaxTemperature = fmt.Sprintf("%.0f", item.Main.TempMax)
		weatherItem.MinTemperature = fmt.Sprintf("%.0f", item.Main.TempMin)
		weatherItem.Id = fmt.Sprintf("%d", item.Weather[0].ID)
		weatherItem.Icon = item.Weather[0].Icon
		weatherItem.Type = item.Weather[0].Main
		forecastResponse.List = append(forecastResponse.List, weatherItem)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(forecastResponse)
}
