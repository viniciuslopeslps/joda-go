package time

type Time interface{}

type time struct{}

func NewTime() Time {
	return &time{}
}
