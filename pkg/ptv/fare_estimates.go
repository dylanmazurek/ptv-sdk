package ptv

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models"
)

// FareEstimate retrieves fare estimates between zones
// minZone: Minimum zone number
// maxZone: Maximum zone number
// request: Additional request parameters
func (c *Client) FareEstimate(minZone, maxZone int, request *models.FareEstimateRequest) (*models.FareEstimateResponse, error) {
	if minZone < 0 {
		return nil, fmt.Errorf("minZone must be non-negative")
	}
	if maxZone < 0 {
		return nil, fmt.Errorf("maxZone must be non-negative")
	}
	if maxZone < minZone {
		return nil, fmt.Errorf("maxZone must be greater than or equal to minZone")
	}

	// Build the URL path
	path := fmt.Sprintf(constants.PATH_V3_FARE_ESTIMATE, minZone, maxZone)

	// Build query parameters
	params := url.Values{}

	if request != nil {
		// Add journey parameter
		if len(request.Journey) > 0 {
			journeyStrs := make([]string, len(request.Journey))
			for i, stopID := range request.Journey {
				journeyStrs[i] = strconv.Itoa(stopID)
			}
			params.Set("journey", strings.Join(journeyStrs, ","))
		}

		// Add is_journey parameter
		if request.IsJourney != nil {
			params.Set("is_journey", strconv.FormatBool(*request.IsJourney))
		}

		// Add travelled parameter
		if request.Travelled != nil {
			params.Set("travelled", strconv.FormatBool(*request.Travelled))
		}
	}
	// Make the API request
	var result models.FareEstimateResponse
	newRequest, err := c.NewRequest("GET", path, nil, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	err = c.Do(newRequest, &result)
	if err != nil {
		return nil, fmt.Errorf("fare estimate request failed: %w", err)
	}

	return &result, nil
}

// GetFareEstimate is a convenience method that retrieves fare estimates between zones
func (c *Client) GetFareEstimate(minZone, maxZone int) (*models.FareEstimateResponse, error) {
	return c.FareEstimate(minZone, maxZone, nil)
}

// GetFareEstimateForJourney retrieves fare estimates for a specific journey between stops
func (c *Client) GetFareEstimateForJourney(originStopID, destinationStopID int, travelled bool) (*models.FareEstimateResponse, error) {
	isJourney := true
	request := &models.FareEstimateRequest{
		Journey:   []int{originStopID, destinationStopID},
		IsJourney: &isJourney,
		Travelled: &travelled,
	}

	// For journey-based estimates, we can use zone 1 to 1 as the base (the API will calculate actual zones)
	return c.FareEstimate(models.Zone1, models.Zone1, request)
}

// GetFareEstimateWithMyki retrieves fare estimates for customers who have used myki
func (c *Client) GetFareEstimateWithMyki(minZone, maxZone int) (*models.FareEstimateResponse, error) {
	travelled := true
	request := &models.FareEstimateRequest{
		Travelled: &travelled,
	}
	return c.FareEstimate(minZone, maxZone, request)
}

// GetFareEstimateWithoutMyki retrieves fare estimates for customers who haven't used myki
func (c *Client) GetFareEstimateWithoutMyki(minZone, maxZone int) (*models.FareEstimateResponse, error) {
	travelled := false
	request := &models.FareEstimateRequest{
		Travelled: &travelled,
	}
	return c.FareEstimate(minZone, maxZone, request)
}

// GetCityCirleFareEstimate retrieves fare estimates for City Circle (free zone)
func (c *Client) GetCityCirleFareEstimate() (*models.FareEstimateResponse, error) {
	return c.GetFareEstimate(models.ZoneCityCircle, models.ZoneCityCircle)
}

// GetZone1FareEstimate retrieves fare estimates for Zone 1 only
func (c *Client) GetZone1FareEstimate() (*models.FareEstimateResponse, error) {
	return c.GetFareEstimate(models.Zone1, models.Zone1)
}

// GetZone2FareEstimate retrieves fare estimates for Zone 2 only
func (c *Client) GetZone2FareEstimate() (*models.FareEstimateResponse, error) {
	return c.GetFareEstimate(models.Zone2, models.Zone2)
}

// GetZone1And2FareEstimate retrieves fare estimates for Zones 1 and 2
func (c *Client) GetZone1And2FareEstimate() (*models.FareEstimateResponse, error) {
	return c.GetFareEstimate(models.Zone1, models.Zone2)
}
