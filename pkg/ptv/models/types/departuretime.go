package types

import (
	"time"
)

type DepartureTime time.Time

func (d *DepartureTime) HumanString() string {
	if d == nil || time.Time(*d).IsZero() {
		return ""
	}

	str := time.Time(*d).Format(time.RFC3339)

	return str
}
