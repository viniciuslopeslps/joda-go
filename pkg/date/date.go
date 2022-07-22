package date

import (
	"fmt"
)

var calendar = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

type Date interface {
	ToString() string
	Sum(value Date) Date
	GetDays() int
	GetMonths() int
	GetYears() int
	IsLeapYear() bool
}

type date struct {
	days   int
	months int
	years  int
}

func NewDate(days, months, years int) Date {
	if (days <= 0 || days > 31) || (months < 0 || months > 12) || (years < 0) {
		panic("invalid date!")
	}

	return &date{
		days:   days,
		months: months,
		years:  years,
	}
}

func (d *date) Sum(value Date) Date {
	days := d.GetDays() + value.GetDays()
	upperBound := d.getMonthBound(d.GetMonths())
	months := d.GetMonths() + value.GetMonths()
	years := d.GetYears() + value.GetYears()

	if days > upperBound {
		diff := days - upperBound
		days = diff
		months += diff % upperBound
	}

	if months > 12 {
		years += months / 12
		months = months % 12
	}

	return &date{
		days:   days,
		months: months,
		years:  years,
	}
}

func (d *date) IsLeapYear() bool {
	years := d.GetYears()
	if years%4 != 0 {
		return false
	}

	if years%100 != 0 {
		return true
	}

	if years%400 == 0 {
		return true
	}
	return false
}

func (d *date) GetDays() int {
	return d.days
}

func (d *date) GetMonths() int {
	return d.months
}

func (d *date) GetYears() int {
	return d.years
}

func (d *date) ToString() string {
	return fmt.Sprintf("%d/%d/%d", d.days, d.months, d.years)
}

func (d *date) getMonthBound(month int) int {
	if month == 2 && d.IsLeapYear() {
		return 29
	}

	return calendar[month]
}
