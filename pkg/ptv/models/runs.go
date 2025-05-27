package models

type Run struct {
	RunID            int    `json:"run_id"`
	RunRef           string `json:"run_ref"`
	RouteID          int    `json:"route_id"`
	RouteType        int    `json:"route_type"`
	FinalStopID      int    `json:"final_stop_id"`
	DestinationName  string `json:"destination_name"`
	Status           string `json:"status"`
	DirectionID      int    `json:"direction_id"`
	RunSequence      int    `json:"run_sequence"`
	ExpressStopCount int    `json:"express_stop_count"`
	RunNote          string `json:"run_note"`
}
