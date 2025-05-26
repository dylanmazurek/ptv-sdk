package ptv

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models"
)

// Outlets retrieves all myki ticket outlets
// request: Additional request parameters and filters
func (c *Client) Outlets(request *models.OutletsRequest) (*models.OutletsResponse, error) {
	// Build the URL path
	path := constants.PATH_V3_OUTLETS

	// Build query parameters
	params := url.Values{}

	if request != nil {
		// Add max results filter
		if request.MaxResults != nil {
			params.Set("max_results", strconv.Itoa(*request.MaxResults))
		}
	}
	// Make the API request
	var result models.OutletsResponse
	newRequest, err := c.NewRequest("GET", path, nil, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	err = c.Do(newRequest, &result)
	if err != nil {
		return nil, fmt.Errorf("outlets request failed: %w", err)
	}

	return &result, nil
}

// OutletsByLocation retrieves myki ticket outlets near a location
// latitude: Geographic coordinate of latitude
// longitude: Geographic coordinate of longitude
// request: Additional request parameters and filters
func (c *Client) OutletsByLocation(latitude, longitude float64, request *models.OutletsByLocationRequest) (*models.OutletsResponse, error) {
	// Build the URL path
	path := fmt.Sprintf(constants.PATH_V3_OUTLETS_BY_LOCATION, latitude, longitude)

	// Build query parameters
	params := url.Values{}

	if request != nil {
		// Add max results filter
		if request.MaxResults != nil {
			params.Set("max_results", strconv.Itoa(*request.MaxResults))
		}

		// Add max distance filter
		if request.MaxDistance != nil {
			params.Set("max_distance", strconv.FormatFloat(*request.MaxDistance, 'f', -1, 64))
		}
	}
	// Make the API request
	var result models.OutletsResponse
	newRequest, err := c.NewRequest("GET", path, nil, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	err = c.Do(newRequest, &result)
	if err != nil {
		return nil, fmt.Errorf("outlets by location request failed: %w", err)
	}

	return &result, nil
}

// GetAllOutlets is a convenience method that retrieves all myki ticket outlets
func (c *Client) GetAllOutlets() (*models.OutletsResponse, error) {
	return c.Outlets(nil)
}

// GetOutletsLimited is a convenience method that retrieves a limited number of outlets
func (c *Client) GetOutletsLimited(maxResults int) (*models.OutletsResponse, error) {
	request := &models.OutletsRequest{
		MaxResults: &maxResults,
	}
	return c.Outlets(request)
}

// GetNearbyOutlets is a convenience method that retrieves outlets near a location
func (c *Client) GetNearbyOutlets(latitude, longitude float64) (*models.OutletsResponse, error) {
	request := &models.OutletsByLocationRequest{
		Latitude:  latitude,
		Longitude: longitude,
	}
	return c.OutletsByLocation(latitude, longitude, request)
}

// GetNearbyOutletsWithin is a convenience method that retrieves outlets within a distance
func (c *Client) GetNearbyOutletsWithin(latitude, longitude, maxDistance float64) (*models.OutletsResponse, error) {
	request := &models.OutletsByLocationRequest{
		Latitude:    latitude,
		Longitude:   longitude,
		MaxDistance: &maxDistance,
	}
	return c.OutletsByLocation(latitude, longitude, request)
}

// GetNearbyOutletsLimited is a convenience method that retrieves a limited number of nearby outlets
func (c *Client) GetNearbyOutletsLimited(latitude, longitude float64, maxResults int, maxDistance *float64) (*models.OutletsResponse, error) {
	request := &models.OutletsByLocationRequest{
		Latitude:    latitude,
		Longitude:   longitude,
		MaxResults:  &maxResults,
		MaxDistance: maxDistance,
	}
	return c.OutletsByLocation(latitude, longitude, request)
}
