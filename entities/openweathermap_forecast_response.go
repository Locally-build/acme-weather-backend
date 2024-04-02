package entities

type OpenWeatherMapForecastApiResponse struct {
	Cod     string               `json:"cod"`
	Message int                  `json:"message"`
	Cnt     int                  `json:"cnt"`
	List    []OpenWeatherMapList `json:"list"`
	City    OpenWeatherMapCity   `json:"city"`
}

type OpenWeatherMapCity struct {
	ID         int                 `json:"id"`
	Name       string              `json:"name"`
	Coord      OpenWeatherMapCoord `json:"coord"`
	Country    string              `json:"country"`
	Population int                 `json:"population"`
	Timezone   int                 `json:"timezone"`
	Sunrise    int                 `json:"sunrise"`
	Sunset     int                 `json:"sunset"`
}

type OpenWeatherMapCoord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type OpenWeatherMapList struct {
	Dt         int                     `json:"dt"`
	Main       OpenWeatherMapMain      `json:"main"`
	Weather    []OpenWeatherMapWeather `json:"weather"`
	Clouds     OpenWeatherMapClouds    `json:"clouds"`
	Wind       OpenWeatherMapWind      `json:"wind"`
	Visibility int                     `json:"visibility"`
	Pop        float64                 `json:"pop"`
	Rain       *OpenWeatherMapRain     `json:"rain,omitempty"`
	Sys        OpenWeatherMapSys       `json:"sys"`
	DtTxt      string                  `json:"dt_txt"`
}

type OpenWeatherMapClouds struct {
	All int `json:"all"`
}

type OpenWeatherMapMain struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
	TempKf    float64 `json:"temp_kf"`
}

type OpenWeatherMapRain struct {
	The1H float64 `json:"1h"`
}

type OpenWeatherMapSys struct {
	Pod string `json:"pod"`
}

type OpenWeatherMapWeather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type OpenWeatherMapWind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}
