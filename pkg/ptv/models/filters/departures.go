package filters

import (
	"net/url"
	"strconv"
	"time"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
)

type DeparturesFilter struct {
	BaseFilter

	RouteType constants.RouteType
	StopID    int
	RouteID   int

	DirectionID      *int
	IncludeCancelled *bool
	AfterDate        *time.Time
	Expand           []constants.ExpandOption
}

func (d *DeparturesFilter) ToURLValues() url.Values {
	params := d.BaseFilter.DefaultValues()

	if d.DirectionID != nil {
		params.Set("direction_id", strconv.Itoa(*d.DirectionID))
	}

	if d.AfterDate != nil {
		params.Set("after_date", d.AfterDate.Format(time.RFC3339))
	}

	if d.IncludeCancelled != nil && *d.IncludeCancelled {
		params.Set("include_cancelled", "true")
	}

	if len(d.Expand) > 0 {
		for _, expand := range d.Expand {
			params.Add("expand", string(expand))
		}
	}

	return params
}
