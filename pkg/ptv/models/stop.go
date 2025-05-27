package models

type Stop struct {
	ID                 int    `json:"stop_id"`
	StationName        string `json:"station_name"`
	StationDescription string `json:"station_description,omitempty"`
	DisruptionIDs      []int  `json:"disruption_ids,omitempty"`
}
