package models

import "encoding/json"

type DisruptionsRequest struct {
	RouteTypes       []int  // Filter by route_type; values returned via RouteTypes API
	DisruptionModes  []int  // Filter by disruption_mode; values returned via v3/disruptions/modes API
	DisruptionStatus string // Filter by status of disruption ("current", "planned")
}

type DisruptionsByRouteRequest struct {
	RouteID          int
	DisruptionStatus string // Filter by status of disruption ("current", "planned")
}

type DisruptionsByStopRequest struct {
	StopID           int
	DisruptionStatus string // Filter by status of disruption ("current", "planned")
}

type DisruptionsByRouteAndStopRequest struct {
	RouteID          int
	StopID           int
	DisruptionStatus string // Filter by status of disruption ("current", "planned")
}

type DisruptionsResponse struct {
	Disruptions map[string]Disruption `json:"disruptions"`
	Status      Status                `json:"status"`
}

type DisruptionModesResponse struct {
	DisruptionModes []DisruptionMode `json:"disruption_modes"`
	Status          Status           `json:"status"`
}

type DisruptionMode struct {
	DisruptionModeID   json.Number `json:"disruption_mode_id"`
	DisruptionModeName string      `json:"disruption_mode_name"`
}

// Enhanced Disruption model (already defined in stops.go but including here for completeness)
type DisruptionComplete struct {
	DisruptionID     json.Number       `json:"disruption_id"`
	Title            string            `json:"title"`
	URL              string            `json:"url"`
	Description      string            `json:"description"`
	DisruptionStatus string            `json:"disruption_status"`
	DisruptionType   string            `json:"disruption_type"`
	PublishedOn      string            `json:"published_on"`
	LastUpdated      string            `json:"last_updated"`
	FromDate         string            `json:"from_date"`
	ToDate           string            `json:"to_date"`
	Routes           []DisruptionRoute `json:"routes"`
	Stops            []DisruptionStop  `json:"stops"`
	Colour           string            `json:"colour"`
	DisplayOnBoard   bool              `json:"display_on_board"`
	DisplayStatus    bool              `json:"display_status"`
}

type DisruptionRouteComplete struct {
	RouteType   json.Number          `json:"route_type"`
	RouteID     json.Number          `json:"route_id"`
	RouteName   string               `json:"route_name"`
	RouteNumber string               `json:"route_number"`
	RouteGTFSID string               `json:"route_gtfs_id"`
	Direction   *DisruptionDirection `json:"direction"`
}

type DisruptionStopComplete struct {
	StopID   json.Number `json:"stop_id"`
	StopName string      `json:"stop_name"`
}

type DisruptionDirectionComplete struct {
	RouteDirectionID json.Number `json:"route_direction_id"`
	DirectionID      json.Number `json:"direction_id"`
	DirectionName    string      `json:"direction_name"`
	ServiceTime      string      `json:"service_time"`
}

// Disruption status constants
const (
	DisruptionStatusCurrent = "current"
	DisruptionStatusPlanned = "planned"
)

// Disruption mode constants (from swagger)
const (
	DisruptionModeGeneralIncident      = 1
	DisruptionModeTrackWork            = 2
	DisruptionModeLineWork             = 3
	DisruptionModeStationWork          = 4
	DisruptionModeStopWork             = 4 // alias for station work
	DisruptionModeEquipmentFailure     = 5
	DisruptionModeIndustrialAction     = 7
	DisruptionModePoliceIncident       = 8
	DisruptionModeEmergencyServices    = 9
	DisruptionModeAccident             = 10
	DisruptionModeWeather              = 11
	DisruptionModeSpecialEvent         = 12
	DisruptionModeSystemFailure        = 13
	DisruptionModeHighPassengerVolumes = 14
	DisruptionModeOther                = 100
)
