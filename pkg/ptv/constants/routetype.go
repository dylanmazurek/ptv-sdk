package constants

type RouteType int

const (
	RouteTypeTrain      RouteType = 0
	RouteTypeTram       RouteType = 1
	RouteTypeBus        RouteType = 2
	RouteTypeVLine      RouteType = 3
	RouteTypeNightRider RouteType = 4
)

func (r RouteType) String() string {
	switch r {
	case RouteTypeTrain:
		return "train"
	case RouteTypeTram:
		return "tram"
	case RouteTypeBus:
		return "bus"
	case RouteTypeVLine:
		return "vline"
	case RouteTypeNightRider:
		return "nightrider"
	default:
		return "unknown"
	}
}

func RouteTypeFromString(s string) (RouteType, bool) {
	switch s {
	case "train":
		return RouteTypeTrain, true
	case "tram":
		return RouteTypeTram, true
	case "bus":
		return RouteTypeBus, true
	case "vline":
		return RouteTypeVLine, true
	case "nightrider":
		return RouteTypeNightRider, true
	default:
		return 0, false
	}
}
