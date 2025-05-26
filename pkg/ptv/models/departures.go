package models

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type DeparturesRequest struct {
	RouteType int
	StopID    int
	RouteID   int

	DirectionID      *int
	GitLookup        *bool
	IncludeCancelled *bool
	AfterDate        *time.Time
	Expand           []ExpandOption

	MaxResults *int
}

func (d *DeparturesRequest) ToURLValues() url.Values {
	params := url.Values{}
	if d.DirectionID != nil {
		params.Set("direction_id", strconv.Itoa(*d.DirectionID))
	}

	if d.GitLookup != nil && *d.GitLookup {
		params.Set("gtfs", "true")
	}

	if d.AfterDate != nil {
		params.Set("after_date", d.AfterDate.Format(time.RFC3339))
	}

	if d.MaxResults != nil {
		params.Set("max_results", strconv.Itoa(*d.MaxResults))
	}

	if d.IncludeCancelled != nil && *d.IncludeCancelled {
		params.Set("include_cancelled", "true")
	}

	if len(d.Expand) > 0 {
		expandStrs := make([]string, len(d.Expand))
		for i, expand := range d.Expand {
			expandStrs[i] = string(expand)
		}
		params.Set("expand", strings.Join(expandStrs, ","))
	}

	return params
}

type Stop struct {
	ID        int    `json:"stop_id"`
	Name      string `json:"stop_name"`
	Distance  int    `json:"stop_distance"`
	Suburb    string `json:"stop_suburb"`
	Latitude  int    `json:"stop_latitude"`
	Longitude int    `json:"stop_longitude"`
	Landmark  string `json:"stop_landmark"`
	Sequence  int    `json:"stop_sequence"`
	RouteType int    `json:"route_type"`
}

type DeparturesResponse struct {
	Departures []Departure `json:"departures"`

	Stops      []Stop      `json:"-"`
	Routes     []Route     `json:"-"`
	Runs       []Run       `json:"-"`
	Directions []Direction `json:"-"`
}

func (d *DeparturesResponse) UnmarshalJSON(data []byte) error {
	type Alias DeparturesResponse
	aux := &struct {
		*Alias

		Stops      map[string]Stop      `json:"stops"`
		Routes     map[string]Route     `json:"routes"`
		Runs       map[string]Run       `json:"runs"`
		Directions map[string]Direction `json:"directions"`
	}{
		Alias: (*Alias)(d),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	return nil
}

type Departure struct {
	StopID            int     `json:"stop_id"`
	RouteID           int     `json:"route_id"`
	RunID             int     `json:"run_id"`
	RunRef            string  `json:"run_ref"`
	DirectionID       int     `json:"direction_id"`
	DisruptionIDs     []int   `json:"disruption_ids"`
	IsAtPlatform      bool    `json:"at_platform"`
	PlatformNumber    *string `json:"platform_number"`
	DepartureSequence int     `json:"departure_sequence"`
	Flags             string  `json:"flags"`

	ScheduledDeparture time.Time  `json:"-"`
	EstimatedDeparture *time.Time `json:"-"`
}

func (d *Departure) UnmarshalJSON(data []byte) error {
	type Alias Departure
	aux := &struct {
		*Alias

		ScheduledDepartureUTC string `json:"scheduled_departure_utc"`
		EstimatedDepartureUTC string `json:"estimated_departure_utc"`
	}{
		Alias: (*Alias)(d),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	if aux.ScheduledDepartureUTC != "" {
		scheduledTime, err := time.Parse(time.RFC3339, aux.ScheduledDepartureUTC)
		if err != nil {
			return err
		}

		d.ScheduledDeparture = scheduledTime
	}

	if aux.EstimatedDepartureUTC != "" {
		estimatedTime, err := time.Parse(time.RFC3339, aux.EstimatedDepartureUTC)
		if err != nil {
			return err
		}

		d.EstimatedDeparture = &estimatedTime
	}

	return nil
}
