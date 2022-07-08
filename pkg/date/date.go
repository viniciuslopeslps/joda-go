package date

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
	Diff(a Date, b Date) Date
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
	return &date{
		days:   days,
		months: months,
		years:  years,
	}
}

func (d *date) Sum(value Date) Date {
	days := d.GetDays() + value.GetDays()
	upperBound := d.getMonthBound(d.GetMonths())

	if days > upperBound {
		diff := days - upperBound

	}
	return nil
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
	//TODO implement me
	panic("implement me")
}

func (d *date) Diff(a Date, b Date) Date {
	//TODO implement me
	panic("implement me")
}

func (d *date) getMonthBound(month int) int {
	if month == 2 && d.IsLeapYear() {
		return 29
	}

	return calendar[month]
}
