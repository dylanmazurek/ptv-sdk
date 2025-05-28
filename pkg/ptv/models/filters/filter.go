package filters

import (
	"net/url"
	"strconv"
)

type BaseFilter struct {
	MaxResults *int `json:"max_results,omitempty"`
}

func (b *BaseFilter) DefaultValues() url.Values {
	params := url.Values{}
	if b.MaxResults != nil {
		params.Set("max_results", strconv.Itoa(*b.MaxResults))
	}

	return params
}
