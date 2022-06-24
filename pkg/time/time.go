package time

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
	return ""
}

func (t *time) Sum(value Time) Time {
	return &time{}
}

func (t *time) Diff(a Time, b Time) Time {
	return &time{}
}
