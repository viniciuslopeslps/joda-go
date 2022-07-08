package localtime

import (
	"fmt"
)

type Time interface {
	ToString() string
	Sum(value Time) Time
	Diff(value Time) Time
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
	if hours > 24 || minutes > 60 || seconds > 60 {
		panic("This is an invalid local time")
	}

	return &Localtime{
		Hours:   hours,
		Minutes: minutes,
		Seconds: seconds,
	}
}

func (t *Localtime) ToString() string {
	return fmt.Sprintf("%.2d:%.2d:%.2d", t.Hours, t.Minutes, t.Seconds)
}

func (t *Localtime) Sum(value Time) Time {
	newSeconds := t.Seconds + value.GetSeconds()
	newMinutes := t.Minutes + value.GetMinutes()
	newHours := t.Hours + value.GetHours()

	newTime := &Localtime{
		Seconds: newSeconds,
		Minutes: newMinutes,
	}
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

func (t *Localtime) Diff(value Time) Time {
	newSeconds := t.Seconds - value.GetSeconds()
	newMinutes := t.Minutes - value.GetMinutes()
	newHours := t.Hours - value.GetHours()
	newTime := &Localtime{}

	if newSeconds < 0 {
		newMinutes += 1
		newTime.Seconds = 60 - newSeconds
	} else {
		newTime.Seconds = newSeconds
	}

	if newMinutes < 0 {
		newHours += 1
		newTime.Minutes = 60 - newMinutes
	} else {
		newTime.Minutes = newMinutes
	}

	newTime.Hours = newHours

	return newTime
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
