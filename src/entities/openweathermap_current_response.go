package entities

type OpenWeatherMapCurrentWeatherResponse struct {
	Coord      OpenWeatherMapCoord     `json:"coord"`
	Weather    []OpenWeatherMapWeather `json:"weather"`
	Base       string                  `json:"base"`
	Main       OpenWeatherMapMain      `json:"main"`
	Visibility int                     `json:"visibility"`
	Wind       OpenWeatherMapWind      `json:"wind"`
	Clouds     OpenWeatherMapClouds    `json:"clouds"`
	Dt         int                     `json:"dt"`
	Sys        OpenWeatherMapSys       `json:"sys"`
	Timezone   int                     `json:"timezone"`
	ID         int                     `json:"id"`
	Name       string                  `json:"name"`
	Cod        int                     `json:"cod"`
}
