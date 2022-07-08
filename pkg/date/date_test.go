package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsLeapYear(t *testing.T) {
	assert.True(t, NewDate(0, 0, 1964).IsLeapYear())
	assert.True(t, NewDate(0, 0, 2000).IsLeapYear())
	assert.False(t, NewDate(0, 0, 1950).IsLeapYear())
	assert.False(t, NewDate(0, 0, 5000).IsLeapYear())
	assert.True(t, NewDate(0, 0, 1032).IsLeapYear())
}
