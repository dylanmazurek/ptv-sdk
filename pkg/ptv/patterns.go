package ptv

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models"
)

// PatternByRun retrieves the stopping pattern for a specific run
// runRef: Identifier of a specific service run
// routeType: Number identifying transport mode
// request: Additional request parameters and filters
func (c *Client) PatternByRun(runRef string, routeType int, request *models.PatternRequest) (*models.PatternResponse, error) {
	if runRef == "" {
		return nil, fmt.Errorf("runRef cannot be empty")
	}
	if routeType < 0 {
		return nil, fmt.Errorf("routeType must be non-negative")
	}

	// Build the URL path
	path := fmt.Sprintf(constants.PATH_V3_PATTERNS_BY_RUN, runRef, routeType)

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

		// Add stop filter
		if request.StopID != nil {
			params.Set("stop_id", strconv.Itoa(*request.StopID))
		}

		// Add date filter
		if request.DateUTC != "" {
			params.Set("date_utc", request.DateUTC)
		}
	}
	// Make the API request
	var result models.PatternResponse
	newRequest, err := c.NewRequest("GET", path, nil, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	err = c.Do(newRequest, &result)
	if err != nil {
		return nil, fmt.Errorf("pattern by run request failed: %w", err)
	}

	return &result, nil
}

// GetPatternByRun is a convenience method that retrieves the stopping pattern for a run
func (c *Client) GetPatternByRun(runRef string, routeType int) (*models.PatternResponse, error) {
	return c.PatternByRun(runRef, routeType, nil)
}

// GetPatternByRunExpanded is a convenience method that retrieves the stopping pattern with expanded data
func (c *Client) GetPatternByRunExpanded(runRef string, routeType int, expand []models.ExpandOption) (*models.PatternResponse, error) {
	request := &models.PatternRequest{
		Expand: expand,
	}
	return c.PatternByRun(runRef, routeType, request)
}

// GetPatternByRunForStop is a convenience method that retrieves the stopping pattern filtered by a specific stop
func (c *Client) GetPatternByRunForStop(runRef string, routeType int, stopID int) (*models.PatternResponse, error) {
	request := &models.PatternRequest{
		StopID: &stopID,
	}
	return c.PatternByRun(runRef, routeType, request)
}

// GetPatternByRunWithDate is a convenience method that retrieves the stopping pattern for a specific date
func (c *Client) GetPatternByRunWithDate(runRef string, routeType int, dateUTC string) (*models.PatternResponse, error) {
	request := &models.PatternRequest{
		DateUTC: dateUTC,
	}
	return c.PatternByRun(runRef, routeType, request)
}
