package models

import "github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"

type Route struct {
	ID            int                 `json:"route_id"`
	RouteType     constants.RouteType `json:"route_type"`
	Name          string              `json:"route_name"`
	Number        string              `json:"route_number"`
	GTFSID        string              `json:"route_gtfs_id"`
	ServiceStatus RouteServiceStatus  `json:"route_service_status"`
}

type RouteServiceStatus struct {
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}
