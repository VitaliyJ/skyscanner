package skyscanner

// Stats object
type Stats struct {
	Itineraries ItineraryStats `json:"itineraries"`
}

// ItineraryStats contains statistics for itineraries
type ItineraryStats struct {
	MinDuration              int32              `json:"minDuration"`
	MaxDuration              int32              `json:"maxDuration"`
	Total                    ItinerarySummary   `json:"total"`
	Stops                    ItineraryStopStats `json:"stops"`
	HasChangeAirportTransfer bool               `json:"hasChangeAirportTransfer"` // Indicates if response contains itineraries with airport transfers
}

// ItinerarySummary summary for a category
type ItinerarySummary struct {
	Count    int32 `json:"count"`
	MinPrice Price `json:"minPrice"`
}

// ItineraryStopStats - itinerary stats based on number of stops
type ItineraryStopStats struct {
	Direct       ItineraryStopSummaryStats `json:"direct"`
	OneStop      ItineraryStopSummaryStats `json:"oneStop"`
	TwoPlusStops ItineraryStopSummaryStats `json:"twoPlusStops"`
}

// ItineraryStopSummaryStats - itinerary stats based on stops
type ItineraryStopSummaryStats struct {
	Total       ItinerarySummary         `json:"total"` // Itinerary summary for a category
	TicketTypes ItineraryStopTicketStats `json:"ticketTypes"`
}

// ItineraryStopTicketStats itinerary stats based on type of ticket
type ItineraryStopTicketStats struct {
	SingleTicket      ItinerarySummary `json:"singleTicket"`
	MultiTicketNonNpt ItinerarySummary `json:"multiTicketNonNpt"`
	MultiTicketNpt    ItinerarySummary `json:"multiTicketNpt"`
}
