package ptv

import (
	"fmt"
	"net/http"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models"
)

func (c *Client) DirectionsForRoute(reqParams models.DirectionsForRouteRequest) ([]models.DirectionWithDescription, error) {
	reqPath := fmt.Sprintf(constants.PATH_V3_DIRECTIONS_FOR_ROUTE, reqParams.RouteID)

	newRequest, err := c.NewRequest(http.MethodGet, reqPath, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var directionsResponse models.DirectionsResponse
	err = c.Do(newRequest, &directionsResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return directionsResponse.Directions, nil
}

func (c *Client) DirectionsByID(reqParams models.DirectionsByIDRequest) ([]models.DirectionWithDescription, error) {
	reqPath := fmt.Sprintf(constants.PATH_V3_DIRECTIONS_BY_ID, reqParams.DirectionID)

	newRequest, err := c.NewRequest(http.MethodGet, reqPath, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var directionsResponse models.DirectionsResponse
	err = c.Do(newRequest, &directionsResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return directionsResponse.Directions, nil
}

func (c *Client) DirectionsByType(reqParams models.DirectionsByTypeRequest) ([]models.DirectionWithDescription, error) {
	reqPath := fmt.Sprintf(constants.PATH_V3_DIRECTIONS_BY_TYPE, reqParams.DirectionID, reqParams.RouteType)

	newRequest, err := c.NewRequest(http.MethodGet, reqPath, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var directionsResponse models.DirectionsResponse
	err = c.Do(newRequest, &directionsResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return directionsResponse.Directions, nil
}
