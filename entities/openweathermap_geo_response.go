package entities

type GeoApiResponse []GeoApiResponseElement

type GeoApiResponseElement struct {
	Name       string            `json:"name"`
	LocalNames map[string]string `json:"local_names"`
	Lat        float64           `json:"lat"`
	Lon        float64           `json:"lon"`
	Country    string            `json:"country"`
	State      *string           `json:"state,omitempty"`
}
