package models

import "encoding/json"

type StopsRequest struct {
	StopID            int
	RouteType         int
	StopLocation      bool // Indicates if stop location information will be returned
	StopAmenities     bool // Indicates if stop amenity information will be returned
	StopAccessibility bool // Indicates if stop accessibility information will be returned
	StopContact       bool // Indicates if stop contact information will be returned
	StopTicket        bool // Indicates if stop ticket information will be returned
	GTFS              bool // Indicates whether the stop_id is a GTFS ID or not
	StopStaffing      bool // Indicates if stop staffing information will be returned
	StopDisruptions   bool // Indicates if stop disruption information will be returned
}

type StopsForRouteRequest struct {
	RouteID                      int
	RouteType                    int
	DirectionID                  *int   // Direction for which the stops need to be returned
	StopDisruptions              bool   // Flag to specify whether disruptions should be included
	IncludeGeopath               bool   // Flag to specify whether geo_path should be included
	GeopathUTC                   string // Filter geopaths by date (ISO 8601 UTC format)
	IncludeAdvertisedInterchange bool   // Flag to specify whether additional stops for interchanges should be included
}

type StopsByLocationRequest struct {
	Latitude        float64
	Longitude       float64
	RouteTypes      []int    // Filter by route_type
	MaxResults      *int     // Maximum number of results returned (default = 30)
	MaxDistance     *float64 // Filter by maximum distance (in metres) from location (default = 300)
	StopDisruptions bool     // Indicates if stop disruption information will be returned
}

type StopResponse struct {
	Stop        StopDetails           `json:"stop"`
	Disruptions map[string]Disruption `json:"disruptions"`
	Status      Status                `json:"status"`
}

type StopsOnRouteResponse struct {
	Stops       []StopOnRoute         `json:"stops"`
	Disruptions map[string]Disruption `json:"disruptions"`
	Geopath     []json.RawMessage     `json:"geopath"`
	Status      Status                `json:"status"`
}

type StopsByDistanceResponse struct {
	Stops       []StopGeosearch       `json:"stops"`
	Disruptions map[string]Disruption `json:"disruptions"`
	Status      Status                `json:"status"`
}

type StopDetails struct {
	DisruptionIDs      []json.Number       `json:"disruption_ids"`
	StationType        string              `json:"station_type"`
	StationDescription string              `json:"station_description"`
	RouteType          json.Number         `json:"route_type"`
	StopLocation       *StopLocation       `json:"stop_location"`
	StopAmenities      *StopAmenityDetails `json:"stop_amenities"`
	StopAccessibility  *StopAccessibility  `json:"stop_accessibility"`
	StopStaffing       *StopStaffing       `json:"stop_staffing"`
	Routes             []json.RawMessage   `json:"routes"`
	StopID             json.Number         `json:"stop_id"`
	StopName           string              `json:"stop_name"`
	StopLandmark       string              `json:"stop_landmark"`
}

type StopModel struct {
	StopDistance  json.Number `json:"stop_distance"`
	StopSuburb    string      `json:"stop_suburb"`
	StopName      string      `json:"stop_name"`
	StopID        json.Number `json:"stop_id"`
	RouteType     json.Number `json:"route_type"`
	StopLatitude  json.Number `json:"stop_latitude"`
	StopLongitude json.Number `json:"stop_longitude"`
	StopLandmark  string      `json:"stop_landmark"`
	StopSequence  json.Number `json:"stop_sequence"`
}

type StopOnRoute struct {
	DisruptionIDs []json.Number      `json:"disruption_ids"`
	StopSuburb    string             `json:"stop_suburb"`
	RouteType     json.Number        `json:"route_type"`
	StopLatitude  json.Number        `json:"stop_latitude"`
	StopLongitude json.Number        `json:"stop_longitude"`
	StopSequence  json.Number        `json:"stop_sequence"`
	StopTicket    *StopTicket        `json:"stop_ticket"`
	Interchange   []InterchangeRoute `json:"interchange"`
	StopID        json.Number        `json:"stop_id"`
	StopName      string             `json:"stop_name"`
	StopLandmark  string             `json:"stop_landmark"`
}

type StopGeosearch struct {
	DisruptionIDs []json.Number     `json:"disruption_ids"`
	StopDistance  json.Number       `json:"stop_distance"`
	StopSuburb    string            `json:"stop_suburb"`
	StopName      string            `json:"stop_name"`
	StopID        json.Number       `json:"stop_id"`
	RouteType     json.Number       `json:"route_type"`
	Routes        []json.RawMessage `json:"routes"`
	StopLatitude  json.Number       `json:"stop_latitude"`
	StopLongitude json.Number       `json:"stop_longitude"`
	StopLandmark  string            `json:"stop_landmark"`
	StopSequence  json.Number       `json:"stop_sequence"`
}

type StopLocation struct {
	GPS StopGPS `json:"gps"`
}

type StopGPS struct {
	Latitude  json.Number `json:"latitude"`
	Longitude json.Number `json:"longitude"`
}

type StopAmenityDetails struct {
	Toilet     bool   `json:"toilet"`
	TaxiRank   bool   `json:"taxi_rank"`
	CarParking string `json:"car_parking"`
	CCTV       bool   `json:"cctv"`
}

type StopAccessibility struct {
	Lighting                      bool                         `json:"lighting"`
	PlatformNumber                json.Number                  `json:"platform_number"`
	AudioCustomerInformation      bool                         `json:"audio_customer_information"`
	Escalator                     bool                         `json:"escalator"`
	HearingLoop                   bool                         `json:"hearing_loop"`
	Lift                          bool                         `json:"lift"`
	Stairs                        bool                         `json:"stairs"`
	StopAccessible                bool                         `json:"stop_accessible"`
	TactileGroundSurfaceIndicator bool                         `json:"tactile_ground_surface_indicator"`
	WaitingRoom                   bool                         `json:"waiting_room"`
	Wheelchair                    *StopAccessibilityWheelchair `json:"wheelchair"`
}

type StopAccessibilityWheelchair struct {
	AccessibleRamp        bool `json:"accessible_ramp"`
	Parking               bool `json:"parking"`
	Telephone             bool `json:"telephone"`
	Toilet                bool `json:"toilet"`
	LowTicketCounter      bool `json:"low_ticket_counter"`
	Manouvering           bool `json:"manouvering"`
	RaisedPlatform        bool `json:"raised_platform"`
	Ramp                  bool `json:"ramp"`
	SecondaryPath         bool `json:"secondary_path"`
	RaisedPlatformShelter bool `json:"raised_platform_shelther"`
	SteepRamp             bool `json:"steep_ramp"`
}

type StopStaffing struct {
	FriAMFrom        string `json:"fri_am_from"`
	FriAMTo          string `json:"fri_am_to"`
	FriPMFrom        string `json:"fri_pm_from"`
	FriPMTo          string `json:"fri_pm_to"`
	MonAMFrom        string `json:"mon_am_from"`
	MonAMTo          string `json:"mon_am_to"`
	MonPMFrom        string `json:"mon_pm_from"`
	MonPMTo          string `json:"mon_pm_to"`
	PHAdditionalText string `json:"ph_additional_text"`
	PHFrom           string `json:"ph_from"`
	PHTo             string `json:"ph_to"`
	SatAMFrom        string `json:"sat_am_from"`
	SatAMTo          string `json:"sat_am_to"`
	SatPMFrom        string `json:"sat_pm_from"`
	SatPMTo          string `json:"sat_pm_to"`
	SunAMFrom        string `json:"sun_am_from"`
	SunAMTo          string `json:"sun_am_to"`
	SunPMFrom        string `json:"sun_pm_from"`
	SunPMTo          string `json:"sun_pm_to"`
	ThuAMFrom        string `json:"thu_am_from"`
	ThuAMTo          string `json:"thu_am_to"`
	ThuPMFrom        string `json:"thu_pm_from"`
	ThuPMTo          string `json:"thu_pm_to"`
	TueAMFrom        string `json:"tue_am_from"`
	TueAMTo          string `json:"tue_am_to"`
	TuePMFrom        string `json:"tue_pm_from"`
	TuePMTo          string `json:"tue_pm_to"`
	WedAMFrom        string `json:"wed_am_from"`
	WedAMTo          string `json:"wed_am_to"`
	WedPMFrom        string `json:"wed_pm_from"`
	WedPMTo          string `json:"wed_pm_To"`
}

type StopTicket struct {
	TicketType       string        `json:"ticket_type"`
	Zone             string        `json:"zone"`
	IsFreeFareZone   bool          `json:"is_free_fare_zone"`
	TicketMachine    bool          `json:"ticket_machine"`
	TicketChecks     bool          `json:"ticket_checks"`
	VlineReservation bool          `json:"vline_reservation"`
	TicketZones      []json.Number `json:"ticket_zones"`
}

type InterchangeRoute struct {
	RouteID    json.Number `json:"route_id"`
	Advertised bool        `json:"advertised"`
}

type Disruption struct {
	DisruptionID     json.Number       `json:"disruption_id"`
	Title            string            `json:"title"`
	URL              string            `json:"url"`
	Description      string            `json:"description"`
	DisruptionStatus string            `json:"disruption_status"`
	DisruptionType   string            `json:"disruption_type"`
	PublishedOn      string            `json:"published_on"`
	LastUpdated      string            `json:"last_updated"`
	FromDate         string            `json:"from_date"`
	ToDate           string            `json:"to_date"`
	Routes           []DisruptionRoute `json:"routes"`
	Stops            []DisruptionStop  `json:"stops"`
	Colour           string            `json:"colour"`
	DisplayOnBoard   bool              `json:"display_on_board"`
	DisplayStatus    bool              `json:"display_status"`
}

type DisruptionRoute struct {
	RouteType   json.Number          `json:"route_type"`
	RouteID     json.Number          `json:"route_id"`
	RouteName   string               `json:"route_name"`
	RouteNumber string               `json:"route_number"`
	RouteGTFSID string               `json:"route_gtfs_id"`
	Direction   *DisruptionDirection `json:"direction"`
}

type DisruptionStop struct {
	StopID   json.Number `json:"stop_id"`
	StopName string      `json:"stop_name"`
}

type DisruptionDirection struct {
	RouteDirectionID json.Number `json:"route_direction_id"`
	DirectionID      json.Number `json:"direction_id"`
	DirectionName    string      `json:"direction_name"`
	ServiceTime      string      `json:"service_time"`
}
