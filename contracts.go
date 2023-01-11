package skyscanner

import "strconv"

const (
	CabinClassUnspecified    CabinClass = "CABIN_CLASS_UNSPECIFIED"
	CabinClassEconomy        CabinClass = "CABIN_CLASS_ECONOMY"
	CabinClassPremiumEconomy CabinClass = "CABIN_CLASS_PREMIUM_ECONOMY"
	CabinClassBusiness       CabinClass = "CABIN_CLASS_BUSINESS"
	CabinClassFirst          CabinClass = "CABIN_CLASS_FIRST"

	ResponseStatusUnspecified ResponseStatus = "RESULT_STATUS_UNSPECIFIED"
	ResponseStatusComplete    ResponseStatus = "RESULT_STATUS_COMPLETE"
	ResponseStatusIncomplete  ResponseStatus = "RESULT_STATUS_INCOMPLETE"
	ResponseStatusFailed      ResponseStatus = "RESULT_STATUS_FAILED"

	ResponseActionUnspecified ResponseAction = "RESULT_ACTION_UNSPECIFIED"
	ResponseActionReplaced    ResponseAction = "RESULT_ACTION_REPLACED"
	ResponseActionNotModified ResponseAction = "RESULT_ACTION_NOT_MODIFIED"
	ResponseActionOmitted     ResponseAction = "RESULT_ACTION_OMITTED"

	PriceUnitUnspecified PriceUnit = "PRICE_UNIT_UNSPECIFIED" // Unit is not specified
	PriceUnitWhole       PriceUnit = "PRICE_UNIT_WHOLE"       // unit relation: 1. eg: A whole pound, euro, etc
	PriceUnitCenti       PriceUnit = "PRICE_UNIT_CENTI"       // unit relation: 100; eg: cents
	PriceUnitMilli       PriceUnit = "PRICE_UNIT_MILLI"       // unit relation: 1000
	PriceUnitMicro       PriceUnit = "PRICE_UNIT_MICRO"       // unit relation: 1000000

	// TransferTypeUnspecified - The transfer type is not specified
	TransferTypeUnspecified TransferType = "TRANSFER_TYPE_UNSPECIFIED"
	// TransferTypeManaged - A protected transfer, managed by the agent.
	// If a flights is delayed and the travellers miss their connection,
	// the agent will find alternatives an no extra cost
	TransferTypeManaged TransferType = "TRANSFER_TYPE_MANAGED"
	// TransferTypeSelfTransfer - A non-protected transfer.
	// The travellers will have multiple booking references.
	// If a flights is delayed and the travellers miss their connection they will have to buy a whole new ticket
	TransferTypeSelfTransfer TransferType = "TRANSFER_TYPE_SELF_TRANSFER"
	// TransferTypeProtectedSelfTransfer - Self transfer flights that are protected by the online travel agent
	// rather than the airline.
	// If a flights is delayed and the travellers miss their connection
	// they can contact the travel agent to help with the rebooking
	TransferTypeProtectedSelfTransfer TransferType = "TRANSFER_TYPE_UNSPECIFIED"

	PlaceTypeUnspecified PlaceType = "PLACE_TYPE_UNSPECIFIED"
	PlaceTypeAirport     PlaceType = "PLACE_TYPE_AIRPORT"
	PlaceTypeCity        PlaceType = "PLACE_TYPE_CITY"
	PlaceTypeCountry     PlaceType = "PLACE_TYPE_COUNTRY"
	PlaceTypeContinent   PlaceType = "PLACE_TYPE_CONTINENT"

	AgentTypeUnspecified AgentType = "AGENT_TYPE_UNSPECIFIED"  // unspecified agent type
	AgentTypeTravelAgent AgentType = "AGENT_TYPE_TRAVEL_AGENT" // agent is a travel agent
	AgentTypeAirline     AgentType = "AGENT_TYPE_AIRLINE"      // agent is a airline
)

type CabinClass string
type ResponseStatus string
type ResponseAction string
type PriceUnit string
type TransferType string
type PlaceType string
type AgentType string

// Client is a SkyScanner client interface
type Client interface {
	Create(req *CreateRequest) (*CreatePollResponse, *ErrorResponse)
	Poll(req *PollRequest) (*CreatePollResponse, *ErrorResponse)
	Locales() (*LocalesResponse, *ErrorResponse)
	Currencies() (*CurrenciesResponse, *ErrorResponse)
	Markets(locale string) (*MarketsResponse, *ErrorResponse)
	NearestCulture(ip string) (*NearestCultureResponse, *ErrorResponse)
	AutoSuggestFlights(req *AutoSuggestFlightsRequest) (*AutoSuggestFlightsResponse, *ErrorResponse)
}

// Price object
type Price struct {
	Amount string    `json:"amount"`
	Unit   PriceUnit `json:"unit"`
}

// LocalDatetime is a object for inputing times without timezone
type LocalDatetime struct {
	// Year in YYYY format. E.g. 2022
	Year int32 `json:"year"`
	// Month in int value. E.g. 2 is January or 10 is October
	Month int32 `json:"month"`
	// Day in int value. E.g. 5 or 28
	Day int32 `json:"day"`
	// Hour in int value. E.g. 3 or 12
	Hour int32 `json:"hour"`
	// Minute in int value. E.g. 2 or 52
	Minute int32 `json:"minute"`
	// Second in int value. E.g. 1 or 46
	Second int32 `json:"second"`
}

func (p Price) ToFloat() (float64, error) {
	if p.Amount == "" {
		return 0, nil
	}

	a, err := strconv.Atoi(p.Amount)
	if err != nil {
		return 0, err
	}

	switch p.Unit {
	case PriceUnitCenti:
		return float64(a) / 100, nil
	case PriceUnitMilli:
		return float64(a) / 1000, nil
	case PriceUnitMicro:
		return float64(a) / 1000000, nil
	case PriceUnitUnspecified:
		fallthrough
	case PriceUnitWhole:
		fallthrough
	default:
		return float64(a), nil
	}
}
