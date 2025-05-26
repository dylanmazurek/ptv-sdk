package models

import "encoding/json"

type RunsForRouteRequest struct {
	RouteID                      int
	Expand                       []ExpandOption // List of objects to be returned in full
	DateUTC                      string         // Filter by the date and time of the request (ISO 8601 UTC format)
	IncludeAdvertisedInterchange bool           // Indicates whether data related to interchanges should be included
}

type RunsForRouteAndTypeRequest struct {
	RouteID                      int
	RouteType                    int
	Expand                       []ExpandOption // List of objects to be returned in full
	DateUTC                      string         // Filter by the date and time of the request (ISO 8601 UTC format)
	IncludeAdvertisedInterchange bool           // Indicates whether data related to interchanges should be included
}

type RunsByRefRequest struct {
	RunRef                       string
	IncludeGeopath               bool           // Indicates if geopath data will be returned
	Expand                       []ExpandOption // List of objects to be returned in full
	DateUTC                      string         // Filter by the date and time of the request (ISO 8601 UTC format)
	IncludeAdvertisedInterchange bool           // Indicates whether data related to interchanges should be included
}

type RunByRefAndTypeRequest struct {
	RunRef         string
	RouteType      int
	Expand         []ExpandOption // List of objects to be returned in full
	DateUTC        string         // Filter by the date and time of the request (ISO 8601 UTC format)
	IncludeGeopath bool           // Indicates if geopath data will be returned
}

type RunsResponse struct {
	Runs   []Run  `json:"runs"`
	Status Status `json:"status"`
}

type RunResponse struct {
	Run    Run    `json:"run"`
	Status Status `json:"status"`
}

type Run struct {
	RunID             json.Number        `json:"run_id"`
	RunRef            string             `json:"run_ref"`
	RouteID           json.Number        `json:"route_id"`
	RouteType         json.Number        `json:"route_type"`
	FinalStopID       json.Number        `json:"final_stop_id"`
	DestinationName   string             `json:"destination_name"`
	Status            string             `json:"status"`
	DirectionID       json.Number        `json:"direction_id"`
	RunSequence       json.Number        `json:"run_sequence"`
	ExpressStopCount  json.Number        `json:"express_stop_count"`
	VehiclePosition   *VehiclePosition   `json:"vehicle_position"`
	VehicleDescriptor *VehicleDescriptor `json:"vehicle_descriptor"`
	Geopath           []json.RawMessage  `json:"geopath"`
	Interchange       *Interchange       `json:"interchange"`
	RunNote           string             `json:"run_note"`
}

type VehiclePosition struct {
	Latitude    json.Number `json:"latitude"`
	Longitude   json.Number `json:"longitude"`
	Easting     json.Number `json:"easting"`
	Northing    json.Number `json:"northing"`
	Direction   string      `json:"direction"`
	Bearing     json.Number `json:"bearing"`
	Supplier    string      `json:"supplier"`
	DatetimeUTC string      `json:"datetime_utc"`
	ExpiryTime  string      `json:"expiry_time"`
}

type VehicleDescriptor struct {
	Operator       string `json:"operator"`
	ID             string `json:"id"`
	LowFloor       bool   `json:"low_floor"`
	AirConditioned bool   `json:"air_conditioned"`
	Description    string `json:"description"`
	Supplier       string `json:"supplier"`
	Length         string `json:"length"`
}

type Interchange struct {
	Feeder      InterchangeRun `json:"feeder"`
	Distributor InterchangeRun `json:"distributor"`
}

type InterchangeRun struct {
	RunRef          string      `json:"run_ref"`
	RouteID         json.Number `json:"route_id"`
	StopID          json.Number `json:"stop_id"`
	Advertised      bool        `json:"advertised"`
	DirectionID     json.Number `json:"direction_id"`
	DestinationName string      `json:"destination_name"`
}
