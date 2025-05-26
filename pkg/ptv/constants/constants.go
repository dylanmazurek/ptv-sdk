package constants

import "time"

const (
	PTV_BASE_URL = "https://timetableapi.ptv.vic.gov.au"
)

const (
	DEFAULT_TIMEZONE    = "Australia/Melbourne"
	DEFAULT_TIMEOUT     = 10 * time.Second
	DEFAULT_MAX_RESULTS = 5
)

const (
	PATH_V3_DEPARTURES             = "/v3/departures/route_type/%d/stop/%d/route/%d"
	PATH_V3_DEPARTURES_FOR_STOP    = "/v3/departures/route_type/%d/stop/%d"
	PATH_V3_ROUTES                 = "/v3/routes"
	PATH_V3_ROUTES_BY_ID           = "/v3/routes/%d"
	PATH_V3_ROUTE_TYPES            = "/v3/route_types"
	PATH_V3_DIRECTIONS_FOR_ROUTE   = "/v3/directions/route/%d"
	PATH_V3_DIRECTIONS_BY_ID       = "/v3/directions/%d"
	PATH_V3_DIRECTIONS_BY_TYPE     = "/v3/directions/%d/route_type/%d"
	PATH_V3_STOPS_BY_ID            = "/v3/stops/%d/route_type/%d"
	PATH_V3_STOPS_FOR_ROUTE        = "/v3/stops/route/%d/route_type/%d"
	PATH_V3_STOPS_BY_LOCATION      = "/v3/stops/location/%f,%f"
	PATH_V3_SEARCH                 = "/v3/search/%s"
	PATH_V3_RUNS_FOR_ROUTE         = "/v3/runs/route/%d"
	PATH_V3_RUNS_FOR_ROUTE_TYPE    = "/v3/runs/route/%d/route_type/%d"
	PATH_V3_PATTERNS_BY_RUN        = "/v3/pattern/run/%s/route_type/%d"
	PATH_V3_DISRUPTIONS            = "/v3/disruptions"
	PATH_V3_DISRUPTIONS_BY_ROUTE   = "/v3/disruptions/route/%d"
	PATH_V3_DISRUPTIONS_BY_STOP    = "/v3/disruptions/stop/%d"
	PATH_V3_DISRUPTIONS_ROUTE_STOP = "/v3/disruptions/route/%d/stop/%d"
	PATH_V3_DISRUPTION_MODES       = "/v3/disruptions/modes"
	PATH_V3_OUTLETS                = "/v3/outlets"
	PATH_V3_OUTLETS_BY_LOCATION    = "/v3/outlets/location/%f,%f"
	PATH_V3_FARE_ESTIMATE          = "/v3/fare_estimate/min_zone/%d/max_zone/%d"
)
