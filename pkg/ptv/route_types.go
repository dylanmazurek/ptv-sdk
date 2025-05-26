package ptv

import (
	"fmt"
	"net/http"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models"
)

func (c *Client) RouteTypes() ([]models.RouteType, error) {
	reqPath := constants.PATH_V3_ROUTE_TYPES

	newRequest, err := c.NewRequest(http.MethodGet, reqPath, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var routeTypesResponse models.RouteTypesResponse
	err = c.Do(newRequest, &routeTypesResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return routeTypesResponse.RouteTypes, nil
}
