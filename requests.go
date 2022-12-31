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
