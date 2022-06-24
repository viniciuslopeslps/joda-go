package time

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

type time struct {
	Hours   int
	Minutes int
	Seconds int
}

func (t *time) ToString() string {
	return fmt.Sprintf("%v:%v:%v", t.Hours, t.Minutes, t.Seconds)
}

func (t *time) Sum(value Time) Time {
	newSeconds := t.Seconds + value.GetSeconds()
	newMinutes := t.Minutes + value.GetMinutes()
	newHours := t.Hours + value.GetSeconds()

	newTime := &time{}
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

func (t *time) Diff(a Time, b Time) Time {
	return &time{}
}

func (t *time) GetSeconds() int {
	return t.Seconds
}

func (t *time) GetMinutes() int {
	return t.Minutes
}

func (t *time) GetHours() int {
	return t.Hours
}
