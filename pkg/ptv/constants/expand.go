package constants

type ExpandOption string

const (
	ExpandAll               ExpandOption = "All"
	ExpandStop              ExpandOption = "Stop"
	ExpandRoute             ExpandOption = "Route"
	ExpandRun               ExpandOption = "Run"
	ExpandDirection         ExpandOption = "Direction"
	ExpandDisruption        ExpandOption = "Disruption"
	ExpandVehicleDescriptor ExpandOption = "VehicleDescriptor"
	ExpandVehiclePosition   ExpandOption = "VehiclePosition"
	ExpandNone              ExpandOption = "None"
)
