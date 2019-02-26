package realtor

// ListingResponse is the results of finding listings
type ListingResponse struct {
	IsSaved       string      `json:"is_saved"`
	SearchID      interface{} `json:"search_id"`
	SeoURL        string      `json:"seo_url"`
	ClientPolygon interface{} `json:"client_polygon"`
	Results       struct {
		Property struct {
			Type  string             `json:"type"`
			Count int                `json:"count"`
			Total int                `json:"total"`
			Items map[string]Listing `json:"items"`
		} `json:"property"`
		LotPrice struct {
			Type  string `json:"type"`
			Items struct {
			} `json:"items"`
		} `json:"lotPrice"`
		Clusters struct {
			Type  string `json:"type"`
			Items struct {
			} `json:"items"`
			Total int `json:"total"`
		} `json:"clusters"`
	} `json:"results"`
}

// Listing is an individual result from a ListingResponse
type Listing struct {
	AgentID         int         `json:"agentId"`
	LdpURL          string      `json:"ldpUrl"`
	ID              string      `json:"id"`
	PropertyID      string      `json:"property_id"`
	SourceID        string      `json:"source_id"`
	PlanID          interface{} `json:"plan_id"`
	SubdivisionID   interface{} `json:"subdivision_id"`
	ListingID       int         `json:"listing_id"`
	IsStatusPending interface{} `json:"isStatusPending"`
	IsSaved         bool        `json:"isSaved"`
	SavedResourceID interface{} `json:"savedResourceId"`
	MprID           int64       `json:"mprId"`
	OfficeName      string      `json:"officeName"`
	Plot            bool        `json:"plot"`
	Position        struct {
		Coordinates []float64 `json:"coordinates"`
		Type        string    `json:"type"`
	} `json:"position"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	ProductType string `json:"product_type"`
	SearchFlags struct {
		IsPriceReduced       bool        `json:"is_price_reduced"`
		IsPending            bool        `json:"is_pending"`
		IsContingent         bool        `json:"is_contingent"`
		IsNewListing         bool        `json:"is_new_listing"`
		IsShortSale          interface{} `json:"is_short_sale"`
		IsShowcase           bool        `json:"is_showcase"`
		HasTour              bool        `json:"has_tour"`
		HasVideo             bool        `json:"has_video"`
		HasRealtorLogo       bool        `json:"has_realtor_logo"`
		IsCobroke            bool        `json:"is_cobroke"`
		IsCoshowProduct      bool        `json:"is_coshow_product"`
		IsForeclosure        bool        `json:"is_foreclosure"`
		IsNewPlan            bool        `json:"is_new_plan"`
		PriceExcludesLand    bool        `json:"price_excludes_land"`
		HasDeals             bool        `json:"has_deals"`
		IsNew                bool        `json:"is_new"`
		IsCostar             bool        `json:"is_costar"`
		IsAptlist            bool        `json:"is_aptlist"`
		IsAddressSuppressed  bool        `json:"is_address_suppressed"`
		IsSuppressMapPin     bool        `json:"is_suppress_map_pin"`
		IsSuppressMap        bool        `json:"is_suppress_map"`
		IsAdvantage          bool        `json:"is_advantage"`
		IsAdvantageBrand     bool        `json:"is_advantage_brand"`
		IsAdvantagePro       bool        `json:"is_advantage_pro"`
		IsEssentials         bool        `json:"is_essentials"`
		IsBasic              bool        `json:"is_basic"`
		IsBasicOptIn         bool        `json:"is_basic_opt_in"`
		IsBasicOptOut        bool        `json:"is_basic_opt_out"`
		HasMatterport        bool        `json:"has_matterport"`
		IsTcpaMessageEnabled bool        `json:"is_tcpa_message_enabled"`
	} `json:"search_flags"`
	OpenHouseDisplay     interface{} `json:"open_house_display"`
	LastSoldPriceDisplay interface{} `json:"last_sold_price_display"`
	LastSoldDateDisplay  interface{} `json:"last_sold_date_display"`
	Sort                 int         `json:"sort"`
	IsCoBroke            bool        `json:"isCoBroke"`
	IsEnhanced           bool        `json:"isEnhanced"`
	IsForeclosure        bool        `json:"isForeclosure"`
	IsPriceReduced       bool        `json:"isPriceReduced"`
	IsAddressSuppressed  bool        `json:"isAddressSuppressed"`
	Bed                  string      `json:"bed"`
	Bath                 string      `json:"bath"`
	LotSize              int         `json:"lotSize"`
	Photo                string      `json:"photo"`
	PhotoCount           int         `json:"photoCount"`
	Price                int         `json:"price"`
	PriceDisplay         string      `json:"priceDisplay"`
	PropertyType         string      `json:"propertyType"`
	Sqft                 interface{} `json:"sqft"`
	SqftDisplay          interface{} `json:"sqft_display"`
	Address              string      `json:"address"`
	City                 string      `json:"city"`
	County               string      `json:"county"`
	State                string      `json:"state"`
	Zip                  string      `json:"zip"`
	FullAddressDisplay   interface{} `json:"full_address_display"`
	PinPrice             interface{} `json:"pinPrice"`
	DataSource           interface{} `json:"data_source"`
	//AgentID              interface{} `json:"agent_id"`
}
