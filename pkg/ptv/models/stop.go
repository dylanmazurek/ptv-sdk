package models

import "github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"

type Stop struct {
	ID                 int                 `json:"stop_id"`
	Name               string              `json:"stop_name"`
	RouteType          constants.RouteType `json:"route_type,omitempty"`
	StationType        string              `json:"station_type,omitempty"`
	StationDescription string              `json:"station_description,omitempty"`
	Landmark           string              `json:"stop_landmark,omitempty"`
	DisruptionIDs      []int               `json:"disruption_ids,omitempty"`
}
