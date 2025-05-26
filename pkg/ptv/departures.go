package ptv

import (
	"fmt"

	"github.com/aws/smithy-go/ptr"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models"
)

func (c *Client) DeparturesForStop(routeTypeID, stopID, routeID int) ([]models.Departure, error) {
	request := &models.DeparturesRequest{
		RouteType: routeTypeID,
		StopID:    stopID,
		RouteID:   routeID,

		MaxResults: ptr.Int(constants.DEFAULT_MAX_RESULTS),
	}

	response, err := c.departures(request)
	if err != nil {
		return nil, fmt.Errorf("failed to get departures for stop: %w", err)
	}

	return response.Departures, nil
}

func (c *Client) departures(request *models.DeparturesRequest) (*models.DeparturesResponse, error) {
	path := fmt.Sprintf(constants.PATH_V3_DEPARTURES, request.RouteType, request.StopID, request.RouteID)
	params := request.ToURLValues()

	var result models.DeparturesResponse
	newRequest, err := c.NewRequest("GET", path, nil, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	err = c.Do(newRequest, &result)
	if err != nil {
		return nil, fmt.Errorf("departures for route request failed: %w", err)
	}

	return &result, nil
}
