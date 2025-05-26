package models

import "encoding/json"

type FareEstimateRequest struct {
	MinZone   int
	MaxZone   int
	Journey   []int // Journey in format of [origin_stop_id, destination_stop_id]
	IsJourney *bool // Indicates whether 'journey' parameter should be interpreted as an origin/destination pair
	Travelled *bool // Indicates whether customer has travelled using a myki card or not
}

type FareEstimateResponse struct {
	FareEstimate *FareEstimate `json:"fare_estimate"`
	Status       Status        `json:"status"`
}

type FareEstimate struct {
	MinFare        json.Number   `json:"min_fare"`
	MaxFare        json.Number   `json:"max_fare"`
	WeekdayPeak    json.Number   `json:"weekday_peak"`
	WeekdayOffPeak json.Number   `json:"weekday_off_peak"`
	WeekendPeak    json.Number   `json:"weekend_peak"`
	WeekendOffPeak json.Number   `json:"weekend_off_peak"`
	HolidayPeak    json.Number   `json:"holiday_peak"`
	HolidayOffPeak json.Number   `json:"holiday_off_peak"`
	TouchOn        FareComponent `json:"touch_on"`
	TouchOff       FareComponent `json:"touch_off"`
	Zones          []Zone        `json:"zones"`
}

type FareComponent struct {
	CityCircle json.Number `json:"city_circle"`
	Zone1      json.Number `json:"zone_1"`
	Zone2      json.Number `json:"zone_2"`
	Zone1And2  json.Number `json:"zone_1_and_2"`
	Zone3      json.Number `json:"zone_3"`
	Zone1To3   json.Number `json:"zone_1_to_3"`
	Zone2And3  json.Number `json:"zone_2_and_3"`
	Zone1To4   json.Number `json:"zone_1_to_4"`
	Zone2To4   json.Number `json:"zone_2_to_4"`
	Zone3And4  json.Number `json:"zone_3_and_4"`
	Zone1To5   json.Number `json:"zone_1_to_5"`
	Zone2To5   json.Number `json:"zone_2_to_5"`
	Zone3To5   json.Number `json:"zone_3_to_5"`
	Zone4And5  json.Number `json:"zone_4_and_5"`
	Zone1To6   json.Number `json:"zone_1_to_6"`
	Zone2To6   json.Number `json:"zone_2_to_6"`
	Zone3To6   json.Number `json:"zone_3_to_6"`
	Zone4To6   json.Number `json:"zone_4_to_6"`
	Zone5And6  json.Number `json:"zone_5_and_6"`
	Zone1To7   json.Number `json:"zone_1_to_7"`
	Zone2To7   json.Number `json:"zone_2_to_7"`
	Zone3To7   json.Number `json:"zone_3_to_7"`
	Zone4To7   json.Number `json:"zone_4_to_7"`
	Zone5To7   json.Number `json:"zone_5_to_7"`
	Zone6And7  json.Number `json:"zone_6_and_7"`
}

type Zone struct {
	ZoneID   json.Number `json:"zone_id"`
	ZoneName string      `json:"zone_name"`
}

// Zone constants for fare calculations
const (
	ZoneCityCircle = 0
	Zone1          = 1
	Zone2          = 2
	Zone3          = 3
	Zone4          = 4
	Zone5          = 5
	Zone6          = 6
	Zone7          = 7
	Zone8          = 8
	Zone9          = 9
	Zone10         = 10
	Zone11         = 11
	Zone12         = 12
	Zone13         = 13
)

// Common zone combinations
const (
	ZoneComboFreeCircle = ZoneCityCircle
	ZoneComboInner      = Zone1
	ZoneComboMiddle     = Zone2
	ZoneComboOuter      = Zone3
)

// Helper methods for fare estimates

// GetFareForTimeOfDay returns the appropriate fare based on time and day type
func (f *FareEstimate) GetFareForTimeOfDay(isWeekend bool, isPeak bool) json.Number {
	if isWeekend {
		if isPeak {
			return f.WeekendPeak
		}
		return f.WeekendOffPeak
	}

	if isPeak {
		return f.WeekdayPeak
	}
	return f.WeekdayOffPeak
}

// GetFareRange returns the min and max fare as a formatted string
func (f *FareEstimate) GetFareRange() (string, string) {
	return f.MinFare.String(), f.MaxFare.String()
}
