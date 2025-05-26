package models

import "encoding/json"

type Status struct {
	Version string      `json:"version"`
	Health  json.Number `json:"health"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  Status `json:"status"`
}

// Common expand options for API requests
type ExpandOption string

const (
	ExpandAll               ExpandOption = "All"
	ExpandStop              ExpandOption = "Stop"
	ExpandRoute             ExpandOption = "Route"
	ExpandRun               ExpandOption = "Run"
	ExpandDirection         ExpandOption = "Direction"
	ExpandDisruption        ExpandOption = "Disruption"
	ExpandVehicleDescriptor ExpandOption = "VehicleDescriptor"
	ExpandVehiclePosition   ExpandOption = "VehiclePosition"
	ExpandNone              ExpandOption = "None"
)

// Route type constants
const (
	RouteTypeTrain      = 0
	RouteTypeTram       = 1
	RouteTypeBus        = 2
	RouteTypeVLine      = 3
	RouteTypeNightRider = 4
)

// Direction reference constants for SIRI
const (
	DirectionIn               = 1
	DirectionOut              = 2
	DirectionUp               = 5
	DirectionDown             = 10
	DirectionClockwise        = 16
	DirectionCounterclockwise = 32
	DirectionInbound          = 65
	DirectionOutbound         = 130
)
