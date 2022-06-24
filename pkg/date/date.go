package date

type Date interface {
	ToString() string
	Sum(value Date) Date
	Diff(a Date, b Date) Date
}
