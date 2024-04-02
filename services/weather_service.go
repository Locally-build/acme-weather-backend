package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Parallels/acme-weather-demo-backend/entities"
	"github.com/go-resty/resty/v2"
)

var globalWeatherService *WeatherService

type WeatherService struct {
	client *resty.Client
	apiKey string
}

func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{
		client: resty.New(),
		apiKey: apiKey,
	}
}

func GetWeatherService() *WeatherService {
	if globalWeatherService == nil {
		globalWeatherService = NewWeatherService(os.Getenv("OPENWEATHERMAP_API_KEY"))
	}
	return globalWeatherService
}

func (s *WeatherService) GetCityCoordinates(city, state, countryCode string) (*entities.GeoApiResponse, error) {
	resp, err := s.client.R().
		SetQueryParams(map[string]string{
			"q":     fmt.Sprintf("%s,%s,%s", city, state, countryCode),
			"appid": s.apiKey,
			"units": "metric",
		}).
		Get("https://api.openweathermap.org/geo/1.0/direct")

	if err != nil {
		return nil, &entities.OpenWeatherMapAPIErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	var apiResponse entities.GeoApiResponse
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

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		apiError := entities.OpenWeatherMapAPIErrorResponse{
			Code:    resp.StatusCode(),
			Message: err.Error(),
		}

		return nil, &apiError
	}

	return &apiResponse, nil
}

func (s *WeatherService) GetCityWeatherForecast(city, state, country string) (*entities.OpenWeatherMapForecastApiResponse, error) {
	client := resty.New()

	geoResponse, err := s.GetCityCoordinates(city, state, country)
	if err != nil {
		return nil, err
	}
	if len(*geoResponse) == 0 {
		return nil, fmt.Errorf("no coordinates found for %s", city)
	}
	geoCityResponse := *geoResponse

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"lat":   fmt.Sprintf("%f", geoCityResponse[0].Lat),
			"lon":   fmt.Sprintf("%f", geoCityResponse[0].Lon),
			"appid": s.apiKey,
			"mode":  "json",
			"units": "metric",
		}).
		Get("https://api.openweathermap.org/data/2.5/forecast")

	if err != nil {
		return nil, err
	}

	var weatherForecastResponse entities.OpenWeatherMapForecastApiResponse
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

	if err := json.Unmarshal(body, &weatherForecastResponse); err != nil {
		apiError := entities.OpenWeatherMapAPIErrorResponse{
			Code:    resp.StatusCode(),
			Message: "unknown error",
		}

		return nil, &apiError
	}

	return &weatherForecastResponse, nil
}
