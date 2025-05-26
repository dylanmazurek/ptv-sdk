package ptv

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models"
)

// Disruptions retrieves all disruptions for all route types
// request: Filter parameters for disruptions
func (c *Client) Disruptions(request *models.DisruptionsRequest) (*models.DisruptionsResponse, error) {
	// Build the URL path
	path := constants.PATH_V3_DISRUPTIONS

	// Build query parameters
	params := url.Values{}

	if request != nil {
		// Add route types filter
		if len(request.RouteTypes) > 0 {
			routeTypesStr := make([]string, len(request.RouteTypes))
			for i, rt := range request.RouteTypes {
				routeTypesStr[i] = strconv.Itoa(rt)
			}
			params.Set("route_types", strings.Join(routeTypesStr, ","))
		}

		// Add disruption modes filter
		if len(request.DisruptionModes) > 0 {
			disruptionModesStr := make([]string, len(request.DisruptionModes))
			for i, dm := range request.DisruptionModes {
				disruptionModesStr[i] = strconv.Itoa(dm)
			}
			params.Set("disruption_modes", strings.Join(disruptionModesStr, ","))
		}

		// Add disruption status filter
		if request.DisruptionStatus != "" {
			params.Set("disruption_status", request.DisruptionStatus)
		}
	}
	// Make the API request
	var result models.DisruptionsResponse
	newRequest, err := c.NewRequest("GET", path, nil, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	err = c.Do(newRequest, &result)
	if err != nil {
		return nil, fmt.Errorf("disruptions request failed: %w", err)
	}

	return &result, nil
}

// DisruptionsByRoute retrieves all disruptions for a particular route
// routeID: Identifier of route
// request: Additional request parameters and filters
func (c *Client) DisruptionsByRoute(routeID int, request *models.DisruptionsByRouteRequest) (*models.DisruptionsResponse, error) {
	if routeID <= 0 {
		return nil, fmt.Errorf("routeID must be greater than 0")
	}

	// Build the URL path
	path := fmt.Sprintf(constants.PATH_V3_DISRUPTIONS_BY_ROUTE, routeID)

	// Build query parameters
	params := url.Values{}

	if request != nil {
		// Add disruption status filter
		if request.DisruptionStatus != "" {
			params.Set("disruption_status", request.DisruptionStatus)
		}
	}
	// Make the API request
	var result models.DisruptionsResponse
	newRequest, err := c.NewRequest("GET", path, nil, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	err = c.Do(newRequest, &result)
	if err != nil {
		return nil, fmt.Errorf("disruptions by route request failed: %w", err)
	}

	return &result, nil
}

// DisruptionsByStop retrieves all disruptions for a particular stop
// stopID: Identifier of stop
// request: Additional request parameters and filters
func (c *Client) DisruptionsByStop(stopID int, request *models.DisruptionsByStopRequest) (*models.DisruptionsResponse, error) {
	if stopID <= 0 {
		return nil, fmt.Errorf("stopID must be greater than 0")
	}

	// Build the URL path
	path := fmt.Sprintf(constants.PATH_V3_DISRUPTIONS_BY_STOP, stopID)

	// Build query parameters
	params := url.Values{}

	if request != nil {
		// Add disruption status filter
		if request.DisruptionStatus != "" {
			params.Set("disruption_status", request.DisruptionStatus)
		}
	}
	// Make the API request
	var result models.DisruptionsResponse
	newRequest, err := c.NewRequest("GET", path, nil, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	err = c.Do(newRequest, &result)
	if err != nil {
		return nil, fmt.Errorf("disruptions by stop request failed: %w", err)
	}

	return &result, nil
}

// DisruptionsByRouteAndStop retrieves all disruptions for a particular route and stop
// routeID: Identifier of route
// stopID: Identifier of stop
// request: Additional request parameters and filters
func (c *Client) DisruptionsByRouteAndStop(routeID, stopID int, request *models.DisruptionsByRouteAndStopRequest) (*models.DisruptionsResponse, error) {
	if routeID <= 0 {
		return nil, fmt.Errorf("routeID must be greater than 0")
	}
	if stopID <= 0 {
		return nil, fmt.Errorf("stopID must be greater than 0")
	}

	// Build the URL path
	path := fmt.Sprintf(constants.PATH_V3_DISRUPTIONS_ROUTE_STOP, routeID, stopID)

	// Build query parameters
	params := url.Values{}

	if request != nil {
		// Add disruption status filter
		if request.DisruptionStatus != "" {
			params.Set("disruption_status", request.DisruptionStatus)
		}
	}
	// Make the API request
	var result models.DisruptionsResponse
	newRequest, err := c.NewRequest("GET", path, nil, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	err = c.Do(newRequest, &result)
	if err != nil {
		return nil, fmt.Errorf("disruptions by route and stop request failed: %w", err)
	}

	return &result, nil
}

// DisruptionModes retrieves all disruption modes
func (c *Client) DisruptionModes() (*models.DisruptionModesResponse, error) {
	// Build the URL path
	path := constants.PATH_V3_DISRUPTION_MODES
	// Make the API request
	var result models.DisruptionModesResponse
	emptyParams := url.Values{}
	newRequest, err := c.NewRequest("GET", path, nil, &emptyParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	err = c.Do(newRequest, &result)
	if err != nil {
		return nil, fmt.Errorf("disruption modes request failed: %w", err)
	}

	return &result, nil
}

// GetAllDisruptions is a convenience method that retrieves all current disruptions
func (c *Client) GetAllDisruptions() (*models.DisruptionsResponse, error) {
	return c.Disruptions(nil)
}

// GetCurrentDisruptions is a convenience method that retrieves all current disruptions
func (c *Client) GetCurrentDisruptions() (*models.DisruptionsResponse, error) {
	request := &models.DisruptionsRequest{
		DisruptionStatus: models.DisruptionStatusCurrent,
	}
	return c.Disruptions(request)
}

// GetPlannedDisruptions is a convenience method that retrieves all planned disruptions
func (c *Client) GetPlannedDisruptions() (*models.DisruptionsResponse, error) {
	request := &models.DisruptionsRequest{
		DisruptionStatus: models.DisruptionStatusPlanned,
	}
	return c.Disruptions(request)
}

// GetDisruptionsByRouteType is a convenience method that retrieves disruptions for specific route types
func (c *Client) GetDisruptionsByRouteType(routeTypes []int) (*models.DisruptionsResponse, error) {
	request := &models.DisruptionsRequest{
		RouteTypes: routeTypes,
	}
	return c.Disruptions(request)
}

// GetCurrentDisruptionsByRoute is a convenience method that retrieves current disruptions for a route
func (c *Client) GetCurrentDisruptionsByRoute(routeID int) (*models.DisruptionsResponse, error) {
	request := &models.DisruptionsByRouteRequest{
		DisruptionStatus: models.DisruptionStatusCurrent,
	}
	return c.DisruptionsByRoute(routeID, request)
}

// GetCurrentDisruptionsByStop is a convenience method that retrieves current disruptions for a stop
func (c *Client) GetCurrentDisruptionsByStop(stopID int) (*models.DisruptionsResponse, error) {
	request := &models.DisruptionsByStopRequest{
		DisruptionStatus: models.DisruptionStatusCurrent,
	}
	return c.DisruptionsByStop(stopID, request)
}
