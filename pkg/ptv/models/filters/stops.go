package filters

import "github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"

type StopsFilter struct {
	StopID     int
	RouteTypes []constants.RouteType

	IncludeLocation      *bool
	IncludeAmenities     *bool
	IncludeAccessibility *bool
	IncludeContact       *bool
	IncludeTicket        *bool
	IncludeStaffing      *bool
	IncludeDisruptions   *bool
}
