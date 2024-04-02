package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Locally-build/acme-weather-backend/controllers"
	"github.com/Locally-build/acme-weather-backend/entities"
	"github.com/Locally-build/acme-weather-backend/services"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var version = "0.0.0"

type WeatherResponse struct {
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func getWeather(city string, apiKey string) (*WeatherResponse, error) {
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"q":     city,
			"appid": apiKey,
			"units": "metric",
		}).
		Get("https://api.openweathermap.org/data/2.5/weather")

	if err != nil {
		return nil, err
	}

	var weatherResponse WeatherResponse
	body := resp.Body()

	if !resp.IsSuccess() {
		var apiError entities.OpenWeatherMapAPIErrorResponse
		if resp.StatusCode() == 401 {
			if err := json.Unmarshal(body, &apiError); err != nil {
				apiError = entities.OpenWeatherMapAPIErrorResponse{
					Code:    resp.StatusCode(),
					Message: err.Error(),
				}
			}
		} else {
			apiError = entities.OpenWeatherMapAPIErrorResponse{
				Code:    resp.StatusCode(),
				Message: string(body),
			}
		}

		return nil, &apiError
	}

	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		apiError := entities.OpenWeatherMapAPIErrorResponse{
			Code:    resp.StatusCode(),
			Message: "unknown error",
		}

		return nil, &apiError
	}

	return &weatherResponse, nil
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	city := params["city"]

	log.Printf("Getting weather for %s", city)
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	weatherResponse, err := getWeather(city, apiKey)
	if err != nil {
		log.Print(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Print(err)
		}
		return
	}

	fmt.Fprintf(w, "The temperature in %s is %.1f°C and the weather is %s.", city, weatherResponse.Main.Temp, weatherResponse.Weather[0].Main)
}

func main() {
	fmt.Printf("********************************************************************************\n")
	fmt.Printf("* Acme Weather Backend API                                                     *\n")
	fmt.Printf("* Version: %s                                                               *\n", version)
	fmt.Printf("********************************************************************************\n")
	services.GetWeatherService()

	router := mux.NewRouter()
	router.HandleFunc("/weather/{city}", weatherHandler).Methods("GET")
	router.HandleFunc("/weather/{country}/{city}/forecast", controllers.CityWeatherForecastHandler).Methods("GET")
	router.HandleFunc("/weather/{country}/{state}/{city}/forecast", controllers.CityWeatherForecastHandler).Methods("GET")
	router.HandleFunc("/geo/{country}/{city}", controllers.GeoLocationHandler).Methods("GET")
	router.HandleFunc("/geo/{country}/{state}/{city}", weatherHandler).Methods("GET")
	router.HandleFunc("/icon/{os}/{icon}", controllers.IconHandler).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS()(router)))
}
