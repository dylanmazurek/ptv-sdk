package models

import (
	"encoding/json"
	"time"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models/types"
)

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

	ScheduledDeparture types.DepartureTime `json:"-"`
	EstimatedDeparture types.DepartureTime `json:"-"`
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

		d.ScheduledDeparture = types.DepartureTime(scheduledTime)
	}

	if aux.EstimatedDepartureUTC != "" {
		estimatedTime, err := time.Parse(time.RFC3339, aux.EstimatedDepartureUTC)
		if err != nil {
			return err
		}

		d.EstimatedDeparture = types.DepartureTime(estimatedTime)
	}

	return nil
}
