package realtor

//https://www.realtor.com/search_result.json

// ListingRequest is the query which finds listings
type ListingRequest struct {
	SearchCriteria string      `json:"search_criteria"`
	City           string      `json:"city"`
	County         string      `json:"county"`
	DiscoveryMode  bool        `json:"discovery_mode"`
	State          string      `json:"state"`
	Postal         interface{} `json:"postal"`
	Sort           interface{} `json:"sort"`
	Position       interface{} `json:"position"`
	Facets         struct {
		BedsMin         string        `json:"beds_min"`
		BedsMax         interface{}   `json:"beds_max"`
		BathsMin        string        `json:"baths_min"`
		BathsMax        interface{}   `json:"baths_max"`
		PriceMax        string        `json:"price_max"`
		PropType        string        `json:"prop_type"`
		SqftMin         interface{}   `json:"sqft_min"`
		SqftMax         interface{}   `json:"sqft_max"`
		AcreMin         int           `json:"acre_min"`
		AcreMax         interface{}   `json:"acre_max"`
		LotUnit         string        `json:"lot_unit"`
		AgeMax          interface{}   `json:"age_max"`
		AgeMin          interface{}   `json:"age_min"`
		Radius          interface{}   `json:"radius"`
		Pets            []interface{} `json:"pets"`
		DaysOnMarket    string        `json:"days_on_market"`
		OpenHouse       interface{}   `json:"open_house"`
		ShowListings    string        `json:"show_listings"`
		Pending         interface{}   `json:"pending"`
		Foreclosure     interface{}   `json:"foreclosure"`
		NewConstruction interface{}   `json:"new_construction"`
		MultiSearch     struct {
		} `json:"multi_search"`
		IncludePendingContingency bool          `json:"include_pending_contingency"`
		FeaturesHash              []interface{} `json:"features_hash"`
	} `json:"facets"`
	SearchController    string        `json:"search_controller"`
	Neighborhood        interface{}   `json:"neighborhood"`
	Street              interface{}   `json:"street"`
	SearchType          string        `json:"searchType"`
	School              interface{}   `json:"school"`
	SchoolDistrict      interface{}   `json:"school_district"`
	University          interface{}   `json:"university"`
	Park                interface{}   `json:"park"`
	Types               []string      `json:"types"`
	SearchFacetsToDTM   []string      `json:"searchFacetsToDTM"`
	SearchFeaturesToDTM []interface{} `json:"searchFeaturesToDTM"`
	PageSize            int           `json:"page_size"`
	ViewportHeight      int           `json:"viewport_height"`
	PinHeight           int           `json:"pin_height"`
	BoundingBox         string        `json:"bounding_box"`
	MaxClusters         int           `json:"max_clusters"`
	MaxPins             int           `json:"max_pins"`
	Pos                 string        `json:"pos"`
}
