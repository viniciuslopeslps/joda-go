package localtime

import (
	"fmt"
)

type Time interface {
	ToString() string
	Sum(value Time) Time
	Diff(a Time, b Time) Time
	GetSeconds() int
	GetMinutes() int
	GetHours() int
}

type Localtime struct {
	Hours   int
	Minutes int
	Seconds int
}

func NewLocalTime(hours, minutes, seconds int) Time {
	fmt.Print("oi")
	return &Localtime{
		Hours:   hours,
		Minutes: minutes,
		Seconds: seconds,
	}
}

func (t *Localtime) ToString() string {
	return fmt.Sprintf("%v:%v:%v", t.Hours, t.Minutes, t.Seconds)
}

func (t *Localtime) Sum(value Time) Time {
	newSeconds := t.Seconds + value.GetSeconds()
	newMinutes := t.Minutes + value.GetMinutes()
	newHours := t.Hours + value.GetSeconds()

	newTime := &Localtime{}
	if newSeconds >= 60 {
		newMinutes += newSeconds / 60
		newTime.Seconds = newSeconds % 60
	}

	if newMinutes >= 60 {
		newHours += newMinutes / 60
		newTime.Minutes = newMinutes % 60
	}

	newTime.Hours = newHours

	return newTime
}

func (t *Localtime) Diff(a Time, b Time) Time {
	return &Localtime{}
}

func (t *Localtime) GetSeconds() int {
	return t.Seconds
}

func (t *Localtime) GetMinutes() int {
	return t.Minutes
}

func (t *Localtime) GetHours() int {
	return t.Hours
}
