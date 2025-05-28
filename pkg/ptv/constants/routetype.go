package constants

type RouteType int

const (
	RouteTypeTrain      RouteType = 0
	RouteTypeTram       RouteType = 1
	RouteTypeBus        RouteType = 2
	RouteTypeVLine      RouteType = 3
	RouteTypeNightRider RouteType = 4
)

var RouteTypeNames = map[RouteType]string{
	RouteTypeTrain:      "train",
	RouteTypeTram:       "tram",
	RouteTypeBus:        "bus",
	RouteTypeVLine:      "vline",
	RouteTypeNightRider: "nightrider",
}

var FriendlyRouteTypeNames = map[RouteType]string{
	RouteTypeTrain:      "Train",
	RouteTypeTram:       "Tram",
	RouteTypeBus:        "Bus",
	RouteTypeVLine:      "V/Line",
	RouteTypeNightRider: "NightRider",
}

func (r RouteType) String() string {
	if name, ok := RouteTypeNames[r]; ok {
		return name
	}

	return "unknown"
}

func RouteTypeFromString(s string) (RouteType, bool) {
	for rt, name := range RouteTypeNames {
		if name == s {
			return rt, true
		}
	}

	return 0, false
}

func (r RouteType) FriendlyString() string {
	if name, ok := FriendlyRouteTypeNames[r]; ok {
		return name
	}

	return "Unknown"
}
