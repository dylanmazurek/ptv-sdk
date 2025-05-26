package models

import "encoding/json"

type DirectionsResponse struct {
	Directions []DirectionWithDescription `json:"directions"`
	Status     Status                     `json:"status"`
}

type DirectionWithDescription struct {
	RouteDirectionDescription string      `json:"route_direction_description"`
	DirectionID               json.Number `json:"direction_id"`
	DirectionName             string      `json:"direction_name"`
	RouteID                   json.Number `json:"route_id"`
	RouteType                 json.Number `json:"route_type"`
}

type Direction struct {
	DirectionID   json.Number `json:"direction_id"`
	DirectionName string      `json:"direction_name"`
	RouteID       json.Number `json:"route_id"`
	RouteType     json.Number `json:"route_type"`
}

type DirectionsForRouteRequest struct {
	RouteID int
}

type DirectionsByIDRequest struct {
	DirectionID int
}

type DirectionsByTypeRequest struct {
	DirectionID int
	RouteType   int
}
