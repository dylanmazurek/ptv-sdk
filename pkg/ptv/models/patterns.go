package models

import "encoding/json"

type PatternRequest struct {
	RunRef    string
	RouteType int
	Expand    []ExpandOption // List of objects to be returned in full
	StopID    *int           // Filter by stop_id
	DateUTC   string         // Filter by the date and time of the request (ISO 8601 UTC format)
}

type PatternResponse struct {
	Departures []Departure `json:"departures"`
	Stops      []Stop      `json:"stops"`
	Routes     []Route     `json:"routes"`
	Runs       []Run       `json:"runs"`
	Directions []Direction `json:"directions"`
	Status     Status      `json:"status"`
}

type PatternDeparture struct {
	StopID                json.Number `json:"stop_id"`
	RouteID               json.Number `json:"route_id"`
	RunID                 json.Number `json:"run_id"`
	RunRef                string      `json:"run_ref"`
	DirectionID           json.Number `json:"direction_id"`
	DisruptionIDs         []int       `json:"disruption_ids"`
	ScheduledDepartureUTC string      `json:"scheduled_departure_utc"`
	EstimatedDepartureUTC *string     `json:"estimated_departure_utc"`
	AtPlatform            bool        `json:"at_platform"`
	PlatformNumber        *string     `json:"platform_number"`
	Flags                 string      `json:"flags"`
	DepartureSequence     json.Number `json:"departure_sequence"`
}

// PatternStop represents a stop in a stopping pattern
type PatternStop struct {
	Stop                  *Stop       `json:"stop"`
	DepartureSequence     json.Number `json:"departure_sequence"`
	SkipStop              bool        `json:"skip_stop"`
	EstimatedDepartureUTC *string     `json:"estimated_departure_utc"`
	ScheduledDepartureUTC string      `json:"scheduled_departure_utc"`
	StopTiming            string      `json:"stop_timing"`
}

// StoppingPattern represents the complete stopping pattern for a run
type StoppingPattern struct {
	RunRef       string        `json:"run_ref"`
	RouteType    json.Number   `json:"route_type"`
	RouteID      json.Number   `json:"route_id"`
	DirectionID  json.Number   `json:"direction_id"`
	Stops        []PatternStop `json:"stops"`
	ExpressStops []int         `json:"express_stops"`
	Geopath      []Geopoint    `json:"geopath"`
}

// Geopoint represents a geographic coordinate in the route path
type Geopoint struct {
	Latitude  json.Number `json:"lat"`
	Longitude json.Number `json:"lon"`
}

// PatternWithStops includes the complete pattern with associated stops, routes, etc.
type PatternWithStops struct {
	Pattern    StoppingPattern `json:"pattern"`
	Stops      []Stop          `json:"stops"`
	Routes     []Route         `json:"routes"`
	Runs       []Run           `json:"runs"`
	Directions []Direction     `json:"directions"`
	Status     Status          `json:"status"`
}
