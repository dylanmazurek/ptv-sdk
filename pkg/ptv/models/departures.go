package models

import (
	"encoding/json"
	"math"
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

	ScheduledDeparture types.DepartureTime  `json:"-"`
	EstimatedDeparture *types.DepartureTime `json:"-"`
}

func (d *Departure) IsDelayed() bool {
	return d.EstimatedDeparture != nil && !time.Time(*d.EstimatedDeparture).IsZero()
}

func (d *Departure) DepartureIsNextDay(timezone *time.Location) bool {
	if d.EstimatedDeparture == nil || time.Time(*d.EstimatedDeparture).IsZero() {
		return false
	}

	scheduledTime := time.Time(d.ScheduledDeparture)
	estimatedTime := time.Time(*d.EstimatedDeparture)

	nowInLocation := time.Now().In(timezone)

	scheduledNextDay := nowInLocation.Day() != scheduledTime.Day() || nowInLocation.Year() != scheduledTime.Year()
	if scheduledNextDay {
		return true
	}

	estimatedNextDay := nowInLocation.Day() != estimatedTime.Day() || nowInLocation.Year() != estimatedTime.Year()

	return estimatedNextDay
}

func (d *Departure) FriendlyDepartureTime(timezone *time.Location) string {
	if !d.IsDelayed() {
		scheduledTimeLocal := time.Time(d.ScheduledDeparture).In(timezone)
		return scheduledTimeLocal.Format("3:04 PM")
	}

	estimatedDepartureLocal := time.Time(*d.EstimatedDeparture).In(timezone)
	return estimatedDepartureLocal.Format("3:04 PM")
}

func (d *Departure) DelayMin() *int {
	if d.IsDelayed() {
		scheduledDeparture := time.Time(d.ScheduledDeparture)
		estimatedDeparture := time.Time(*d.EstimatedDeparture)

		estimatedDelay := estimatedDeparture.Sub(scheduledDeparture)
		if estimatedDelay > 0 {
			delayMinutes := int(math.RoundToEven(estimatedDelay.Minutes()))
			if delayMinutes > 0 {
				return &delayMinutes
			}
		}
	}

	return nil
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

		departureTime := types.DepartureTime(estimatedTime)
		d.EstimatedDeparture = &departureTime
	}

	return nil
}
