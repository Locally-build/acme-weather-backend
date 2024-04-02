package entities

type ForecastApiResponse struct {
	City string    `json:"city"`
	List []Weather `json:"list"`
}
