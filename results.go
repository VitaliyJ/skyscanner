package skyscanner

// ErrorResponse contains error response data
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// CreatePollResponse contains Create response data
type CreatePollResponse struct {
	SessionToken string         `json:"sessionToken"`
	Status       ResponseStatus `json:"status"`
	Action       ResponseAction `json:"action"`
	Content      *Content       `json:"content"`
}

// Content Search content object containing results and metadata
type Content struct {
	Results        *Results        `json:"results"`
	Stats          *Stats          `json:"stats"`
	SortingOptions *SortingOptions `json:"sortingOptions"`
}

// Results contains search results object.
type Results struct {
	Itineraries map[string]ItineraryResult `json:"itineraries"`
	Legs        map[string]FlightLeg       `json:"legs"`
	Segments    map[string]Segment         `json:"segments"`
	Places      map[string]Place           `json:"places"`
	Carriers    map[string]Carrier         `json:"carriers"`
	Agents      map[string]Agent           `json:"agents"`
	Alliances   map[string]Alliance        `json:"alliances"`
}

// ItineraryResult contains data for itinerary
type ItineraryResult struct {
	PricingOptions     []PricingOption    `json:"pricingOptions"`
	LegIds             []string           `json:"legIds"`
	SustainabilityData SustainabilityData `json:"sustainabilityData"`
}

// PricingOption is a pricing option for an itinerary
type PricingOption struct {
	Price        Price                   `json:"price"`
	AgentIds     []string                `json:"agentIds"`
	Items        []LivePricingOptionItem `json:"items"`
	TransferType TransferType            `json:"transferType"`
}

// LivePricingOptionItem contains pricing item data
type LivePricingOptionItem struct {
	Price    Price                        `json:"price"`
	AgentID  string                       `json:"agentId"`
	DeepLink string                       `json:"deepLink"`
	Fares    []LivePricingOptionItemFares `json:"fares"`
}

// LivePricingOptionItemFares contains fare data for option item
type LivePricingOptionItemFares struct {
	SegmentID     string `json:"segmentId"`
	BookingCode   string `json:"bookingCode"`
	FareBasisCode string `json:"fareBasisCode"`
}

// SustainabilityData is a Sustainability data object
type SustainabilityData struct {
	IsEcoContender    bool    `json:"isEcoContender"`
	EcoContenderDelta float32 `json:"ecoContenderDelta"`
}

// FlightLeg contains legs search results data
type FlightLeg struct {
	OriginPlaceID       string        `json:"originPlaceId"`
	DestinationPlaceID  string        `json:"destinationPlaceId"`
	DepartureDateTime   LocalDatetime `json:"departureDateTime"`
	ArrivalDateTime     LocalDatetime `json:"arrivalDateTime"`
	DurationInMinutes   int32         `json:"durationInMinutes"`
	StopCount           int32         `json:"stopCount"`
	MarketingCarrierIds []string      `json:"marketingCarrierIds"`
	OperatingCarrierIds []string      `json:"operatingCarrierIds"`
	SegmentIds          []string      `json:"segmentIds"`
}

// Segment contains data for segment
type Segment struct {
	OriginPlaceID         string        `json:"originPlaceId"`
	DestinationPlaceID    string        `json:"destinationPlaceId"`
	DepartureDateTime     LocalDatetime `json:"departureDateTime"`
	ArrivalDateTime       LocalDatetime `json:"arrivalDateTime"`
	DurationInMinutes     int32         `json:"durationInMinutes"`
	MarketingFlightNumber string        `json:"marketingFlightNumber"`
	MarketingCarrierIds   []string      `json:"marketingCarrierIds"`
	OperatingCarrierIds   []string      `json:"operatingCarrierIds"`
}

// Place contains the place details
type Place struct {
	// EntityId is a unique ID for the place.
	// It's an internal ID and it doesn't have any meaning outside of SkyScanner APIs
	EntityId string `json:"entityId"`
	// ParentId is a reference to another place, i.e. an airport can have a parent place which is a city
	ParentId string `json:"parentId"`
	// Name is a localised name of the place
	Name string `json:"name"`
	// Type of the place
	Type PlaceType `json:"type"`
	// Iata - The IATA code of the place. It will only be set for airports and cities
	IATA string `json:"iata"`
}

// Carrier contains data for carrier
type Carrier struct {
	Name       string `json:"name"`
	AllianceID string `json:"allianceId"`
	ImageURL   string `json:"imageUrl"`
	IATA       string `json:"iata"`
}

// Agent contains data for agent
type Agent struct {
	Name                 string               `json:"name"`
	Type                 AgentType            `json:"type"`
	ImageURL             string               `json:"imageUrl"`
	FeedbackCount        int32                `json:"feedbackCount"`        // Number of users who gave feedback
	Rating               float32              `json:"rating"`               // Agent rating
	RatingBreakdown      AgentRatingBreakdown `json:"ratingBreakdown"`      // Agent rating split by different criteria
	IsOptimisedForMobile bool                 `json:"isOptimisedForMobile"` // Indicates if partner website is optimized for mobile
}

// AgentRatingBreakdown contains agent rating split by different criteria
type AgentRatingBreakdown struct {
	CustomerService float32 `json:"customerService"`
	ReliablePrices  float32 `json:"reliablePrices"`
	ClearExtraFees  float32 `json:"clearExtraFees"`
	EaseOfBooking   float32 `json:"easeOfBooking"`
	Other           float32 `json:"other"`
}

// Alliance details.
// An airline alliance is an aviation industry arrangement between two
// or more airlines agreeing to cooperate on a substantial level.
// For example Star Alliance
type Alliance struct {
	Name string `json:"name"`
}

// LocalesResponse contains locales query response
type LocalesResponse struct {
	Status  ResponseStatus `json:"status"`
	Locales []Locale       `json:"locales"`
}

// Locale contains locale data
type Locale struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// CurrenciesResponse contains currency query response
type CurrenciesResponse struct {
	Status     ResponseStatus `json:"status"`
	Currencies []Currency     `json:"currencies"`
}

// Currency contains currency format information
type Currency struct {
	Code                        string `json:"code"`
	Symbol                      string `json:"symbol"`
	ThousandsSeparator          string `json:"thousandsSeparator"`
	DecimalSeparator            string `json:"decimalSeparator"`
	SymbolOnLeft                bool   `json:"symbolOnLeft"`
	SpaceBetweenAmountAndSymbol bool   `json:"spaceBetweenAmountAndSymbol"`
	DecimalDigits               int32  `json:"decimalDigits"`
}

// MarketsResponse contains market query response
type MarketsResponse struct {
	Status  ResponseStatus `json:"status"`
	Markets []Market       `json:"markets"`
}

// Market contains market data
type Market struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// NearestCultureResponse contains nearestculture query response
type NearestCultureResponse struct {
	Status   ResponseStatus `json:"status"`
	Market   Market         `json:"market"`
	Locale   Locale         `json:"locale"`
	Currency Currency       `json:"currency"`
}
