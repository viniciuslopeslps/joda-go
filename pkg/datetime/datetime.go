package datetime

import (
	"joda-go/pkg/date"
	"joda-go/pkg/localtime"
)

type Datetime interface {
	Crop() interface{}
	ToString() string
}

type dateTime struct {
	date date.Date
	time localtime.Time
}

func (d dateTime) ToString() string {
	//TODO implement me
	panic("implement me")
}

func (d dateTime) Crop() interface{} {
	//TODO implement me
	panic("implement me")
}

func NewDateTime(day, month, year, hour, minute, second int) Datetime {
	return &dateTime{}
}
