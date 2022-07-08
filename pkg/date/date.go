package date

type Date interface {
	ToString() string
	Sum(value Date) Date
	Diff(a Date, b Date) Date
	GetDays() int
	GetMonths() int
	GetYears() int
}

type date struct {
	days   int
	months int
	years  int
}

func (d *date) Sum(value Date) Date {
	newDays := 
}

func (d *date) IsLeapYear(value Date) bool{
	if value % 
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
