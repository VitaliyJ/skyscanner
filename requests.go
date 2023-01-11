package skyscanner

// CreateRequest contains Create request attributes
type CreateRequest struct {
	Query *CreateRequestQuery `json:"query"`
}

// CreateRequestQuery contains Create query attributes
type CreateRequestQuery struct {
	// required fields
	Market    string      `json:"market"`
	Locale    string      `json:"locale"`
	Currency  string      `json:"currency"`
	QueryLegs []*QueryLeg `json:"query_legs"`
	Adults    int32       `json:"adults"`

	// required for us
	IncludedCarriersIds []string `json:"includedCarriersIds,omitempty"`

	// optional fields
	CabinClass                CabinClass `json:"cabinClass,omitempty"`
	ChildrenAges              []int      `json:"childrenAges,omitempty"`
	ExcludedCarriersIds       []string   `json:"excludedCarriersIds,omitempty"`
	IncludedAgentsIds         []string   `json:"includedAgentsIds,omitempty"`
	ExcludedAgentsIds         []string   `json:"excludedAgentsIds,omitempty"`
	IncludeSustainabilityData bool       `json:"includeSustainabilityData,omitempty"`
	NearbyAirports            bool       `json:"nearbyAirports,omitempty"`
}

// QueryLeg contains leg data to search for.
type QueryLeg struct {
	OriginPlaceId      *PlaceID       `json:"originPlaceId"`
	DestinationPlaceId *PlaceID       `json:"destinationPlaceId,omitempty"`
	Date               *LocalDatetime `json:"date"`
}

// PlaceID is a place query object.
type PlaceID struct {
	IATA     string        `json:"iata,omitempty"`
	EntityId string        `json:"entityId,omitempty"`
	Date     LocalDatetime `json:"date,omitempty"`
}

// PollRequest contains Poll request attributes
type PollRequest struct {
	SessionToken string `json:"sessionToken"`
}

// AutoSuggestFlightsRequest contains autosuggest/flights request data
type AutoSuggestFlightsRequest struct {
	// Query Object containing query parameters for flight autosuggest search.
	Query AutoSuggestFlightsRequestQuery `json:"query"`
	// Limit Limits number of entities returned in response. Takes a minimum value of 1 and a maximum of 50.
	Limit AutoSuggestFlightsRequestLimit `json:"limit"`
	// IsDestination Alters ranking logic of entities.
	IsDestination bool `json:"isDestination"`
}

// AutoSuggestFlightsRequestQuery contains autosuggest/flights request query data
type AutoSuggestFlightsRequestQuery struct {
	// Locale that the results are returned in. e.g. en-GB
	Locale string `json:"locale"`
	// Market for which the search is for. e.g. UK
	Market string `json:"market"`
	// SearchTerm Term to get autosuggest results for. Omitting the searchTerm will return the most popular destinations.
	SearchTerm string `json:"searchTerm"`
	// IncludedEntityTypes Items Enum: "PLACE_TYPE_AIRPORT" "PLACE_TYPE_CITY" "PLACE_TYPE_COUNTRY"
	// List of entity types to be returned. If empty, all entity types will be returned
	IncludedEntityTypes []PlaceType `json:"includedEntityTypes"`
}

// AutoSuggestFlightsRequestLimit contains autosuggest/flights request limit data
type AutoSuggestFlightsRequestLimit struct {
	Empty   bool  `json:"empty"`
	Present bool  `json:"present"`
	AsInt   int32 `json:"asInt"`
}
