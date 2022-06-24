package time

import "fmt"

type Time interface {
	ToString() string
	Sum(value Time) Time
	Diff(a Time, b Time) Time
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
	return &time{}
}

func (t *time) Diff(a Time, b Time) Time {
	return &time{}
}
