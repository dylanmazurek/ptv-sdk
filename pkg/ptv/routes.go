package ptv

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models"
)

func (c *Client) Routes(reqParams models.RoutesRequest) ([]models.Route, error) {
	reqPath := constants.PATH_V3_ROUTES

	var reqQuery = url.Values{}

	// Add route_types filter if provided
	if len(reqParams.RouteTypes) > 0 {
		for _, routeType := range reqParams.RouteTypes {
			reqQuery.Add("route_types", strconv.Itoa(routeType))
		}
	}

	// Add route_name filter if provided
	if reqParams.RouteName != "" {
		reqQuery.Add("route_name", reqParams.RouteName)
	}

	newRequest, err := c.NewRequest(http.MethodGet, reqPath, nil, &reqQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var routesResponse models.RoutesResponse
	err = c.Do(newRequest, &routesResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return routesResponse.Routes, nil
}

func (c *Client) RouteByID(reqParams models.RouteByIDRequest) (*models.RouteWithStatus, error) {
	reqPath := fmt.Sprintf(constants.PATH_V3_ROUTES_BY_ID, reqParams.RouteID)

	var reqQuery = url.Values{}

	// Add include_geopath parameter if specified
	if reqParams.IncludeGeopath {
		reqQuery.Add("include_geopath", "true")
	}

	// Add geopath_utc parameter if provided
	if reqParams.GeopathUTC != "" {
		reqQuery.Add("geopath_utc", reqParams.GeopathUTC)
	}

	newRequest, err := c.NewRequest(http.MethodGet, reqPath, nil, &reqQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var routeResponse models.RouteResponse
	err = c.Do(newRequest, &routeResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return &routeResponse.Route, nil
}
