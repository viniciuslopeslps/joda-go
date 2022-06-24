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
	if newSeconds >= 60 {
		//min := int(newSeconds / 60)
		difSec := newSeconds % 60
		newSeconds = difSec
	}
	//newMinutes := t.Minutes + value.GetMinutes()
	//newHours := t.Hours + value.GetSeconds()

	return &time{}
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
