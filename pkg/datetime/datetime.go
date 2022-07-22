package datetime

import (
	"fmt"
	"joda-go/pkg/date"
	"joda-go/pkg/localtime"
)

type Datetime interface {
	CropDate() localtime.Time
	CropTime() date.Date
	ToString() string
}

type dateTime struct {
	date date.Date
	time localtime.Time
}

func (d dateTime) ToString() string {
	return fmt.Sprintf("%s %s", d.date.ToString(), d.time.ToString())
}

func (d dateTime) CropDate() localtime.Time {
	return d.time
}

func (d dateTime) CropTime() date.Date {
	return d.date
}

func NewDateTime(day, month, year, hour, minute, second int) Datetime {
	return &dateTime{
		date: date.NewDate(day, month, year),
		time: localtime.NewLocalTime(hour, minute, second),
	}
}
