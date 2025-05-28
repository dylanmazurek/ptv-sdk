package ptv

import (
	"time"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/constants"
)

type Options struct {
	BaseURL   string
	UserID    string
	AccessKey string
	Timezone  *time.Location
}

func DefaultOptions() Options {
	defaultTimezone, err := time.LoadLocation(constants.DEFAULT_TIMEZONE)
	if err != nil {
		defaultTimezone, _ = time.LoadLocation("UTC")
	}

	defaultOptions := Options{
		Timezone: defaultTimezone,
		BaseURL:  constants.PTV_BASE_URL,
	}

	return defaultOptions
}

type Option func(*Options)

func WithBaseURL(baseURL string) Option {
	return func(opts *Options) {
		opts.BaseURL = baseURL
	}
}

func WithUserID(userId string) Option {
	return func(opts *Options) {
		opts.UserID = userId
	}
}

func WithAccessKey(key string) Option {
	return func(opts *Options) {
		opts.AccessKey = key
	}
}

func WithTimezone(timezone *time.Location) Option {
	return func(opts *Options) {
		opts.Timezone = timezone
	}
}
