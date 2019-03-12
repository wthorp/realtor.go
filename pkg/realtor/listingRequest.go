package realtor

//https://www.realtor.com/search_result.json

// ListingRequest is the query which finds listings
type ListingRequest struct {
	SearchCriteria      string   `json:"search_criteria,omitempty"`
	City                string   `json:"city,omitempty"`
	County              string   `json:"county,omitempty"`
	DiscoveryMode       bool     `json:"discovery_mode,omitempty"`
	State               string   `json:"state,omitempty"`
	Postal              int      `json:"postal,omitempty"`
	Sort                string   `json:"sort,omitempty"`
	Position            string   `json:"position,omitempty"`
	Facets              Facets   `json:"facets,omitempty"`
	SearchController    string   `json:"search_controller,omitempty"`
	Neighborhood        string   `json:"neighborhood,omitempty"`
	Street              string   `json:"street,omitempty"`
	SearchType          string   `json:"searchType,omitempty"`
	School              string   `json:"school,omitempty"`
	SchoolDistrict      string   `json:"school_district,omitempty"`
	University          string   `json:"university,omitempty"`
	Park                string   `json:"park,omitempty"`
	Types               []string `json:"types,omitempty"`
	SearchFacetsToDTM   []string `json:"searchFacetsToDTM,omitempty"`
	SearchFeaturesToDTM []string `json:"searchFeaturesToDTM,omitempty"`
	PageSize            int      `json:"page_size,omitempty"`
	ViewportHeight      int      `json:"viewport_height,omitempty"`
	PinHeight           int      `json:"pin_height,omitempty"`
	BoundingBox         string   `json:"bounding_box,omitempty"`
	MaxClusters         int      `json:"max_clusters,omitempty"`
	MaxPins             int      `json:"max_pins,omitempty"`
	Pos                 string   `json:"pos,omitempty"`
}

// Facets is the query facets which finds listings
type Facets struct {
	BedsMin         int      `json:"beds_min,omitempty"`
	BedsMax         int      `json:"beds_max,omitempty"`
	BathsMin        int      `json:"baths_min,omitempty"`
	BathsMax        int      `json:"baths_max,omitempty"`
	PriceMax        int      `json:"price_max,omitempty"`
	PropType        string   `json:"prop_type,omitempty"`
	SqftMin         int      `json:"sqft_min,omitempty"`
	SqftMax         int      `json:"sqft_max,omitempty"`
	AcreMin         int      `json:"acre_min,omitempty"`
	AcreMax         int      `json:"acre_max,omitempty"`
	LotUnit         string   `json:"lot_unit,omitempty"`
	AgeMax          int      `json:"age_max,omitempty"`
	AgeMin          int      `json:"age_min,omitempty"`
	Radius          string   `json:"radius,omitempty"`
	Pets            []string `json:"pets,omitempty"`
	DaysOnMarket    string   `json:"days_on_market,omitempty"`
	OpenHouse       bool     `json:"open_house,omitempty"`
	ShowListings    string   `json:"show_listings,omitempty"`
	Pending         bool     `json:"pending,omitempty"`
	Foreclosure     bool     `json:"foreclosure,omitempty"`
	NewConstruction bool     `json:"new_construction,omitempty"`
	MultiSearch     struct {
	} `json:"multi_search,omitempty"`
	IncludePendingContingency bool     `json:"include_pending_contingency,omitempty"`
	FeaturesHash              []string `json:"features_hash,omitempty"`
}
