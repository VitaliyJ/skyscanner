package skyscanner

// SortingOptions contains data for sorting by best, cheapest or fastest criteria
type SortingOptions struct {
	Best     []SortingOptionItem `json:"best"`
	Cheapest []SortingOptionItem `json:"cheapest"`
	Fastest  []SortingOptionItem `json:"fastest"`
}

// SortingOptionItem contains sorting option for a criteria
type SortingOptionItem struct {
	Score       float32 `json:"score"`
	ItineraryID string  `json:"itineraryId"`
}
