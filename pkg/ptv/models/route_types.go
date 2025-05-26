package models

type RouteTypesResponse struct {
	RouteTypes []RouteType `json:"route_types"`
	Status     Status      `json:"status"`
}

type RouteType struct {
	RouteType     int    `json:"route_type"`
	RouteTypeName string `json:"route_type_name"`
}
