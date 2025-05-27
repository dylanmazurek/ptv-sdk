package ptv

import (
	"fmt"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models"
	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models/filters"
)

func (c *Client) Departures(f *filters.DeparturesFilter) (*models.Response, error) {
	path := fmt.Sprintf(constants.PATH_V3_DEPARTURES, f.RouteType, f.StopID, f.RouteID)
	params := f.ToURLValues()

	var result models.Response
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
