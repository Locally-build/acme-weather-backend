package entities

type Weather struct {
	DateTime       string `json:"dateTime"`
	Id             string `json:"id"`
	Type           string `json:"type"`
	Temperature    string `json:"temperature"`
	Description    string `json:"description"`
	FeelsLike      string `json:"feelsLike"`
	Humidity       string `json:"humidity"`
	WindSpeed      string `json:"windSpeed"`
	MaxTemperature string `json:"maxTemperature"`
	MinTemperature string `json:"minTemperature"`
	Pressure       string `json:"pressure"`
	Icon           string `json:"icon"`
}
