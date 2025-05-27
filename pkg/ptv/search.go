package ptv

// // Search performs a general search across stops, routes, and outlets
// // searchTerm: Search term (stations, routes, addresses, etc.)
// // request: Search options and filters
// func (c *Client) Search(searchTerm string, request *models.SearchRequest) (*models.SearchResult, error) {
// 	if searchTerm == "" {
// 		return nil, fmt.Errorf("search term cannot be empty")
// 	}

// 	// URL encode the search term
// 	encodedTerm := url.QueryEscape(searchTerm)

// 	// Build the URL path
// 	path := fmt.Sprintf(constants.PATH_V3_SEARCH, encodedTerm)

// 	params := url.Values{}

// 	if request != nil {
// 		if len(request.RouteTypes) > 0 {
// 			routeTypesStr := make([]string, len(request.RouteTypes))
// 			for i, rt := range request.RouteTypes {
// 				routeTypesStr[i] = strconv.Itoa(rt)
// 			}
// 			params.Set("route_types", strings.Join(routeTypesStr, ","))
// 		}

// 		if request.Latitude != nil {
// 			params.Set("latitude", strconv.FormatFloat(*request.Latitude, 'f', -1, 64))
// 		}

// 		if request.Longitude != nil {
// 			params.Set("longitude", strconv.FormatFloat(*request.Longitude, 'f', -1, 64))
// 		}

// 		if request.MaxDistance != nil {
// 			params.Set("max_distance", strconv.FormatFloat(*request.MaxDistance, 'f', -1, 64))
// 		}

// 		// Add boolean parameters (only set if true to avoid unnecessary query params)
// 		if request.IncludeAddresses {
// 			params.Set("include_addresses", "true")
// 		}

// 		if !request.IncludeOutlets {
// 			params.Set("include_outlets", "false")
// 		}

// 		if !request.MatchStopBySuburb {
// 			params.Set("match_stop_by_suburb", "false")
// 		}

// 		if !request.MatchRouteBySuburb {
// 			params.Set("match_route_by_suburb", "false")
// 		}

// 		if request.MatchStopByGTFSStopID {
// 			params.Set("match_stop_by_gtfs_stop_id", "true")
// 		}
// 	}
// 	// Make the API request
// 	var result models.SearchResult
// 	newRequest, err := c.NewRequest("GET", path, nil, &params)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create request: %w", err)
// 	}

// 	err = c.Do(newRequest, &result)
// 	if err != nil {
// 		return nil, fmt.Errorf("search request failed: %w", err)
// 	}

// 	return &result, nil
// }
