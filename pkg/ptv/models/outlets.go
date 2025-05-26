package models

import "encoding/json"

type OutletsRequest struct {
	MaxResults *int // Maximum number of results returned
}

type OutletsByLocationRequest struct {
	Latitude    float64
	Longitude   float64
	MaxResults  *int     // Maximum number of results returned
	MaxDistance *float64 // Filter by maximum distance (in metres) from location
}

type OutletsResponse struct {
	Outlets []Outlet `json:"outlets"`
	Status  Status   `json:"status"`
}

type Outlet struct {
	OutletSlidSpid         string       `json:"outlet_slid_spid"`
	OutletName             string       `json:"outlet_name"`
	OutletBusiness         string       `json:"outlet_business"`
	OutletLatitude         json.Number  `json:"outlet_latitude"`
	OutletLongitude        json.Number  `json:"outlet_longitude"`
	OutletSuburb           string       `json:"outlet_suburb"`
	OutletPostcode         json.Number  `json:"outlet_postcode"`
	OutletBusinessHourMon  string       `json:"outlet_business_hour_mon"`
	OutletBusinessHourTue  string       `json:"outlet_business_hour_tue"`
	OutletBusinessHourWed  string       `json:"outlet_business_hour_wed"`
	OutletBusinessHourThur string       `json:"outlet_business_hour_thur"`
	OutletBusinessHourFri  string       `json:"outlet_business_hour_fri"`
	OutletBusinessHourSat  string       `json:"outlet_business_hour_sat"`
	OutletBusinessHourSun  string       `json:"outlet_business_hour_sun"`
	OutletNotes            string       `json:"outlet_notes"`
	OutletDistance         *json.Number `json:"outlet_distance,omitempty"`
}

// OutletWithDistance includes distance information for location-based searches
type OutletWithDistance struct {
	Outlet
	OutletDistance json.Number `json:"outlet_distance"`
}

// OutletBusinessHours represents the operating hours for an outlet
type OutletBusinessHours struct {
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`
	Sunday    string `json:"sunday"`
}

// GetBusinessHours returns a structured representation of business hours
func (o *Outlet) GetBusinessHours() OutletBusinessHours {
	return OutletBusinessHours{
		Monday:    o.OutletBusinessHourMon,
		Tuesday:   o.OutletBusinessHourTue,
		Wednesday: o.OutletBusinessHourWed,
		Thursday:  o.OutletBusinessHourThur,
		Friday:    o.OutletBusinessHourFri,
		Saturday:  o.OutletBusinessHourSat,
		Sunday:    o.OutletBusinessHourSun,
	}
}

// IsOpenToday checks if the outlet has hours listed for a given day
// day should be "monday", "tuesday", etc. (lowercase)
func (o *Outlet) IsOpenToday(day string) bool {
	hours := o.GetBusinessHours()
	switch day {
	case "monday":
		return hours.Monday != "" && hours.Monday != "Closed"
	case "tuesday":
		return hours.Tuesday != "" && hours.Tuesday != "Closed"
	case "wednesday":
		return hours.Wednesday != "" && hours.Wednesday != "Closed"
	case "thursday":
		return hours.Thursday != "" && hours.Thursday != "Closed"
	case "friday":
		return hours.Friday != "" && hours.Friday != "Closed"
	case "saturday":
		return hours.Saturday != "" && hours.Saturday != "Closed"
	case "sunday":
		return hours.Sunday != "" && hours.Sunday != "Closed"
	default:
		return false
	}
}
