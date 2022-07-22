package date

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLeapYear(t *testing.T) {
	assert.True(t, NewDate(10, 10, 1964).IsLeapYear())
	assert.True(t, NewDate(1, 12, 2000).IsLeapYear())
	assert.False(t, NewDate(10, 12, 1950).IsLeapYear())
	assert.False(t, NewDate(10, 12, 5000).IsLeapYear())
	assert.True(t, NewDate(10, 12, 1032).IsLeapYear())
}

func TestSumDate_1(t *testing.T) {
	original := NewDate(1, 1, 1964)
	summed := original.Sum(NewDate(1, 1, 1))
	assert.Equal(t, summed.GetYears(), 1965)
	assert.Equal(t, summed.GetDays(), 2)
	assert.Equal(t, summed.GetMonths(), 2)
}

func TestSumDate_2(t *testing.T) {
	original := NewDate(1, 12, 1964)
	summed := original.Sum(NewDate(1, 1, 0))
	assert.Equal(t, 1965, summed.GetYears())
	assert.Equal(t, 2, summed.GetDays())
	assert.Equal(t, 1, summed.GetMonths())
}

func TestSumDate_3(t *testing.T) {
	original := NewDate(31, 12, 1964)
	summed := original.Sum(NewDate(30, 2, 0))
	assert.Equal(t, 1965, summed.GetYears())
	assert.Equal(t, 30, summed.GetDays())
	assert.Equal(t, 2, summed.GetMonths())
}

func TestSumDate_4(t *testing.T) {
	original := NewDate(31, 7, 1964)
	summed := original.Sum(NewDate(1, 0, 0))
	assert.Equal(t, 1964, summed.GetYears())
	assert.Equal(t, 1, summed.GetDays())
	assert.Equal(t, 8, summed.GetMonths())
}

func TestToString(t *testing.T) {
	original := NewDate(31, 12, 1964)
	assert.Equal(t, "31/12/1964", original.ToString())
}

func TestToString_2(t *testing.T) {
	original := NewDate(31, 1, 1964)
	assert.Equal(t, "31/01/1964", original.ToString())
}
