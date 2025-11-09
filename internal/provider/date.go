package provider

import "time"

type DatetimeProvider struct {
}

func NewDatetimeProvider() *DatetimeProvider {
	return &DatetimeProvider{}
}

func (*DatetimeProvider) Now() time.Time {
	return time.Now().UTC()
}
