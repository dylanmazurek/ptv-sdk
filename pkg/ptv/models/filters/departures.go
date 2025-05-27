package filters

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
)

type DeparturesFilter struct {
	Filter

	RouteType constants.RouteType `schema:"route_type,required"`
	StopID    int                 `schema:"stop_id,required"`
	RouteID   int                 `schema:"route_id,required"`

	DirectionID      *int
	IncludeCancelled *bool
	AfterDate        *time.Time
	Expand           []constants.ExpandOption
}

func (d *DeparturesFilter) ToURLValues() url.Values {
	params := url.Values{}
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
		expandStrs := make([]string, len(d.Expand))
		for i, expand := range d.Expand {
			expandStrs[i] = string(expand)
		}

		params.Set("expand", strings.Join(expandStrs, ","))
	}

	return params
}
