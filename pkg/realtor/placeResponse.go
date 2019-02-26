package realtor

//PlaceResponse is the result of finding places by name
type PlaceResponse struct {
	Result []Place `json:"result"`
}

//Place is a single result from PlaceResponse
type Place struct {
	AreaType  string  `json:"area_type"`
	ID        string  `json:"_id"`
	Score     float64 `json:"_score"`
	City      string  `json:"city"`
	StateCode string  `json:"state_code"`
	Country   string  `json:"country"`
	Centroid  struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"centroid"`
}
