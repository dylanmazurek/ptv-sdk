package models

import "encoding/json"

type SearchRequest struct {
	SearchTerm            string
	RouteTypes            []int    // Filter by route_type
	Latitude              *float64 // Filter by geographic coordinate of latitude
	Longitude             *float64 // Filter by geographic coordinate of longitude
	MaxDistance           *float64 // Filter by maximum distance (in metres) from location
	IncludeAddresses      bool     // Placeholder for future development; currently unavailable
	IncludeOutlets        bool     // Indicates if outlets will be returned in response (default = true)
	MatchStopBySuburb     bool     // Indicates whether to find stops by suburbs in the search term (default = true)
	MatchRouteBySuburb    bool     // Indicates whether to find routes by suburbs in the search term (default = true)
	MatchStopByGTFSStopID bool     // Indicates whether to search for stops according to a metlink stop ID (default = false)
}

type SearchResult struct {
	Stops   []ResultStop   `json:"stops"`
	Routes  []ResultRoute  `json:"routes"`
	Outlets []ResultOutlet `json:"outlets"`
	Status  Status         `json:"status"`
}

type ResultStop struct {
	StopDistance  json.Number   `json:"stop_distance"`
	StopSuburb    string        `json:"stop_suburb"`
	RouteType     json.Number   `json:"route_type"`
	Routes        []ResultRoute `json:"routes"`
	StopLatitude  json.Number   `json:"stop_latitude"`
	StopLongitude json.Number   `json:"stop_longitude"`
	StopSequence  json.Number   `json:"stop_sequence"`
	StopID        json.Number   `json:"stop_id"`
	StopName      string        `json:"stop_name"`
	StopLandmark  string        `json:"stop_landmark"`
}

type ResultRoute struct {
	RouteName          string              `json:"route_name"`
	RouteNumber        string              `json:"route_number"`
	RouteType          json.Number         `json:"route_type"`
	RouteID            json.Number         `json:"route_id"`
	RouteGTFSID        string              `json:"route_gtfs_id"`
	RouteServiceStatus *RouteServiceStatus `json:"route_service_status"`
}

type ResultOutlet struct {
	OutletDistance         json.Number `json:"outlet_distance"`
	OutletSlidSpid         string      `json:"outlet_slid_spid"`
	OutletName             string      `json:"outlet_name"`
	OutletBusiness         string      `json:"outlet_business"`
	OutletLatitude         json.Number `json:"outlet_latitude"`
	OutletLongitude        json.Number `json:"outlet_longitude"`
	OutletSuburb           string      `json:"outlet_suburb"`
	OutletPostcode         json.Number `json:"outlet_postcode"`
	OutletBusinessHourMon  string      `json:"outlet_business_hour_mon"`
	OutletBusinessHourTue  string      `json:"outlet_business_hour_tue"`
	OutletBusinessHourWed  string      `json:"outlet_business_hour_wed"`
	OutletBusinessHourThur string      `json:"outlet_business_hour_thur"`
	OutletBusinessHourFri  string      `json:"outlet_business_hour_fri"`
	OutletBusinessHourSat  string      `json:"outlet_business_hour_sat"`
	OutletBusinessHourSun  string      `json:"outlet_business_hour_sun"`
	OutletNotes            string      `json:"outlet_notes"`
}
