package models

type Direction struct {
	ID   int    `json:"direction_id"`
	Name string `json:"direction_name"`

	DirectionDescription string `json:"route_direction_description"`

	RouteID   int `json:"route_id"`
	RouteType int `json:"route_type"`
}
