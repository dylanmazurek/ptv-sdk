package ptv

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models"
)

func (c *Client) StopDetails(reqParams models.StopsRequest) (*models.StopDetails, error) {
	reqPath := fmt.Sprintf(constants.PATH_V3_STOPS_BY_ID, reqParams.StopID, reqParams.RouteType)

	var reqQuery = url.Values{}

	// Add optional parameters
	if reqParams.StopLocation {
		reqQuery.Add("stop_location", "true")
	}
	if reqParams.StopAmenities {
		reqQuery.Add("stop_amenities", "true")
	}
	if reqParams.StopAccessibility {
		reqQuery.Add("stop_accessibility", "true")
	}
	if reqParams.StopContact {
		reqQuery.Add("stop_contact", "true")
	}
	if reqParams.StopTicket {
		reqQuery.Add("stop_ticket", "true")
	}
	if reqParams.GTFS {
		reqQuery.Add("gtfs", "true")
	}
	if reqParams.StopStaffing {
		reqQuery.Add("stop_staffing", "true")
	}
	if reqParams.StopDisruptions {
		reqQuery.Add("stop_disruptions", "true")
	}

	newRequest, err := c.NewRequest(http.MethodGet, reqPath, nil, &reqQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var stopResponse models.StopResponse
	err = c.Do(newRequest, &stopResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return &stopResponse.Stop, nil
}

func (c *Client) StopsForRoute(reqParams models.StopsForRouteRequest) ([]models.StopOnRoute, error) {
	reqPath := fmt.Sprintf(constants.PATH_V3_STOPS_FOR_ROUTE, reqParams.RouteID, reqParams.RouteType)

	var reqQuery = url.Values{}

	// Add optional parameters
	if reqParams.DirectionID != nil {
		reqQuery.Add("direction_id", strconv.Itoa(*reqParams.DirectionID))
	}
	if reqParams.StopDisruptions {
		reqQuery.Add("stop_disruptions", "true")
	}
	if reqParams.IncludeGeopath {
		reqQuery.Add("include_geopath", "true")
	}
	if reqParams.GeopathUTC != "" {
		reqQuery.Add("geopath_utc", reqParams.GeopathUTC)
	}
	if reqParams.IncludeAdvertisedInterchange {
		reqQuery.Add("include_advertised_interchange", "true")
	}

	newRequest, err := c.NewRequest(http.MethodGet, reqPath, nil, &reqQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var stopsResponse models.StopsOnRouteResponse
	err = c.Do(newRequest, &stopsResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return stopsResponse.Stops, nil
}

func (c *Client) StopsByLocation(reqParams models.StopsByLocationRequest) ([]models.StopGeosearch, error) {
	reqPath := fmt.Sprintf(constants.PATH_V3_STOPS_BY_LOCATION, reqParams.Latitude, reqParams.Longitude)

	var reqQuery = url.Values{}

	// Add route_types filter if provided
	if len(reqParams.RouteTypes) > 0 {
		for _, routeType := range reqParams.RouteTypes {
			reqQuery.Add("route_types", strconv.Itoa(routeType))
		}
	}

	// Add max_results if provided
	if reqParams.MaxResults != nil {
		reqQuery.Add("max_results", strconv.Itoa(*reqParams.MaxResults))
	}

	// Add max_distance if provided
	if reqParams.MaxDistance != nil {
		reqQuery.Add("max_distance", fmt.Sprintf("%.2f", *reqParams.MaxDistance))
	}

	// Add stop_disruptions if specified
	if reqParams.StopDisruptions {
		reqQuery.Add("stop_disruptions", "true")
	}

	newRequest, err := c.NewRequest(http.MethodGet, reqPath, nil, &reqQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var stopsResponse models.StopsByDistanceResponse
	err = c.Do(newRequest, &stopsResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return stopsResponse.Stops, nil
}
