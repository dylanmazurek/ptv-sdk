package models

import "encoding/json"

type RoutesRequest struct {
	RouteTypes []int  // Filter by route_type
	RouteName  string // Filter by name of route (accepts partial matches)
}

type RoutesResponse struct {
	Routes []Route `json:"routes"`
	Status Status  `json:"status"`
}

type RouteResponse struct {
	Route  RouteWithStatus `json:"route"`
	Status Status          `json:"status"`
}

type RouteWithStatus struct {
	RouteServiceStatus RouteServiceStatus `json:"route_service_status"`
	RouteType          json.Number        `json:"route_type"`
	RouteID            json.Number        `json:"route_id"`
	RouteName          string             `json:"route_name"`
	RouteNumber        string             `json:"route_number"`
	RouteGTFSID        string             `json:"route_gtfs_id"`
	Geopath            []json.RawMessage  `json:"geopath"`
}

type Route struct {
	RouteType   json.Number       `json:"route_type"`
	RouteID     json.Number       `json:"route_id"`
	RouteName   string            `json:"route_name"`
	RouteNumber string            `json:"route_number"`
	RouteGTFSID string            `json:"route_gtfs_id"`
	Geopath     []json.RawMessage `json:"geopath"`
}

type RouteServiceStatus struct {
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}

type RouteByIDRequest struct {
	RouteID        int
	IncludeGeopath bool   // Indicates if geopath data will be returned
	GeopathUTC     string // Filter geopaths by date (ISO 8601 UTC format)
}
