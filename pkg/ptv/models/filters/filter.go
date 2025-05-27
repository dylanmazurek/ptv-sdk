package filters

type Filter interface {
	ToURLValues() map[string]string
}

type BaseFilter struct {
	MaxResults *int `json:"max_results,omitempty"`
}
