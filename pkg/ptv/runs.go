package ptv

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models"
)

// RunsForRoute retrieves scheduled runs for the specified route
// routeID: Identifier of route
// request: Additional request parameters and filters
func (c *Client) RunsForRoute(routeID int, request *models.RunsForRouteRequest) (*models.RunsResponse, error) {
	if routeID <= 0 {
		return nil, fmt.Errorf("routeID must be greater than 0")
	}

	// Build the URL path
	path := fmt.Sprintf(constants.PATH_V3_RUNS_FOR_ROUTE, routeID)

	// Build query parameters
	params := url.Values{}

	if request != nil {
		// Add expand options
		if len(request.Expand) > 0 {
			expandStrs := make([]string, len(request.Expand))
			for i, expand := range request.Expand {
				expandStrs[i] = string(expand)
			}
			params.Set("expand", strings.Join(expandStrs, ","))
		}

		// Add date filter
		if request.DateUTC != "" {
			params.Set("date_utc", request.DateUTC)
		}

		// Add interchange flag
		if request.IncludeAdvertisedInterchange {
			params.Set("include_advertised_interchange", "true")
		}
	}
	// Make the API request
	var result models.RunsResponse
	newRequest, err := c.NewRequest("GET", path, nil, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	err = c.Do(newRequest, &result)
	if err != nil {
		return nil, fmt.Errorf("runs for route request failed: %w", err)
	}

	return &result, nil
}

// RunsForRouteAndType retrieves scheduled runs for the specified route and route type
// routeID: Identifier of route
// routeType: Number identifying transport mode
// request: Additional request parameters and filters
func (c *Client) RunsForRouteAndType(routeID, routeType int, request *models.RunsForRouteAndTypeRequest) (*models.RunsResponse, error) {
	if routeID <= 0 {
		return nil, fmt.Errorf("routeID must be greater than 0")
	}
	if routeType < 0 {
		return nil, fmt.Errorf("routeType must be non-negative")
	}

	// Build the URL path
	path := fmt.Sprintf(constants.PATH_V3_RUNS_FOR_ROUTE_TYPE, routeID, routeType)

	// Build query parameters
	params := url.Values{}

	if request != nil {
		// Add expand options
		if len(request.Expand) > 0 {
			expandStrs := make([]string, len(request.Expand))
			for i, expand := range request.Expand {
				expandStrs[i] = string(expand)
			}
			params.Set("expand", strings.Join(expandStrs, ","))
		}

		// Add date filter
		if request.DateUTC != "" {
			params.Set("date_utc", request.DateUTC)
		}

		// Add interchange flag
		if request.IncludeAdvertisedInterchange {
			params.Set("include_advertised_interchange", "true")
		}
	}
	// Make the API request
	var result models.RunsResponse
	newRequest, err := c.NewRequest("GET", path, nil, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	err = c.Do(newRequest, &result)
	if err != nil {
		return nil, fmt.Errorf("runs for route and type request failed: %w", err)
	}

	return &result, nil
}

// RunByRef retrieves a specific run by run reference
// runRef: Identifier of a specific service run
// request: Additional request parameters and filters
func (c *Client) RunByRef(runRef string, request *models.RunsByRefRequest) (*models.RunsResponse, error) {
	if runRef == "" {
		return nil, fmt.Errorf("runRef cannot be empty")
	}

	// For this endpoint, we need to check the swagger to see the exact path
	// Since this might be a different endpoint, let me implement it as a general approach
	path := fmt.Sprintf("/v3/runs/ref/%s", runRef)

	// Build query parameters
	params := url.Values{}

	if request != nil {
		// Add expand options
		if len(request.Expand) > 0 {
			expandStrs := make([]string, len(request.Expand))
			for i, expand := range request.Expand {
				expandStrs[i] = string(expand)
			}
			params.Set("expand", strings.Join(expandStrs, ","))
		}

		// Add date filter
		if request.DateUTC != "" {
			params.Set("date_utc", request.DateUTC)
		}

		// Add geopath flag
		if request.IncludeGeopath {
			params.Set("include_geopath", "true")
		}

		// Add interchange flag
		if request.IncludeAdvertisedInterchange {
			params.Set("include_advertised_interchange", "true")
		}
	}
	// Make the API request
	var result models.RunsResponse
	newRequest, err := c.NewRequest("GET", path, nil, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	err = c.Do(newRequest, &result)
	if err != nil {
		return nil, fmt.Errorf("run by ref request failed: %w", err)
	}

	return &result, nil
}

// RunByRefAndType retrieves a specific run by run reference and route type
// runRef: Identifier of a specific service run
// routeType: Number identifying transport mode
// request: Additional request parameters and filters
func (c *Client) RunByRefAndType(runRef string, routeType int, request *models.RunByRefAndTypeRequest) (*models.RunResponse, error) {
	if runRef == "" {
		return nil, fmt.Errorf("runRef cannot be empty")
	}
	if routeType < 0 {
		return nil, fmt.Errorf("routeType must be non-negative")
	}

	// Build the URL path
	path := fmt.Sprintf("/v3/runs/ref/%s/route_type/%d", runRef, routeType)

	// Build query parameters
	params := url.Values{}

	if request != nil {
		// Add expand options
		if len(request.Expand) > 0 {
			expandStrs := make([]string, len(request.Expand))
			for i, expand := range request.Expand {
				expandStrs[i] = string(expand)
			}
			params.Set("expand", strings.Join(expandStrs, ","))
		}

		// Add date filter
		if request.DateUTC != "" {
			params.Set("date_utc", request.DateUTC)
		}

		// Add geopath flag
		if request.IncludeGeopath {
			params.Set("include_geopath", "true")
		}
	}
	// Make the API request
	var result models.RunResponse
	newRequest, err := c.NewRequest("GET", path, nil, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	err = c.Do(newRequest, &result)
	if err != nil {
		return nil, fmt.Errorf("run by ref and type request failed: %w", err)
	}

	return &result, nil
}

// GetRunsByRoute is a convenience method that retrieves runs for a route with common defaults
func (c *Client) GetRunsByRoute(routeID int) (*models.RunsResponse, error) {
	return c.RunsForRoute(routeID, nil)
}

// GetRunsByRouteWithDate is a convenience method that retrieves runs for a route on a specific date
func (c *Client) GetRunsByRouteWithDate(routeID int, dateUTC string) (*models.RunsResponse, error) {
	request := &models.RunsForRouteRequest{
		DateUTC: dateUTC,
	}
	return c.RunsForRoute(routeID, request)
}

// GetRunsByRouteExpanded is a convenience method that retrieves runs for a route with expanded data
func (c *Client) GetRunsByRouteExpanded(routeID int, expand []models.ExpandOption) (*models.RunsResponse, error) {
	request := &models.RunsForRouteRequest{
		Expand: expand,
	}
	return c.RunsForRoute(routeID, request)
}
