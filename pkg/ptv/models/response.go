package models

import "encoding/json"

type Response struct {
	Departures  []Departure  `json:"-"`
	Stops       []Stop       `json:"-"`
	Routes      []Route      `json:"-"`
	Disruptions []Disruption `json:"-"`

	Status Status `json:"status"`
}

func (d *Response) UnmarshalJSON(data []byte) error {
	type Alias Response
	aux := &struct {
		*Alias

		Departures  json.RawMessage `json:"departures"`
		Stops       json.RawMessage `json:"stops"`
		Routes      json.RawMessage `json:"routes"`
		Disruptions json.RawMessage `json:"disruptions"`
	}{
		Alias: (*Alias)(d),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	if string(aux.Departures) != "{}" {
		var departures []Departure
		err := json.Unmarshal(aux.Departures, &departures)
		if err != nil {
			return err
		}

		d.Departures = departures
	}

	if string(aux.Stops) != "{}" {
		var stops map[string]Stop
		err := json.Unmarshal(aux.Stops, &stops)
		if err != nil {
			return err
		}

		for _, stop := range stops {
			d.Stops = append(d.Stops, stop)
		}
	}

	if string(aux.Routes) != "{}" {
		var routes map[string]Route
		err := json.Unmarshal(aux.Routes, &routes)
		if err != nil {
			return err
		}

		for _, route := range routes {
			d.Routes = append(d.Routes, route)
		}
	}

	if string(aux.Disruptions) != "{}" {
		var disruptions []Disruption
		err := json.Unmarshal(aux.Disruptions, &disruptions)
		if err != nil {
			return err
		}

		d.Disruptions = disruptions
	}

	return nil
}
