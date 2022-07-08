package localtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sum(t *testing.T) {
	testTime := NewLocalTime(1, 1, 1)
	value := NewLocalTime(1, 1, 1)
	expected := NewLocalTime(2, 2, 2)

	res := testTime.Sum(value)
	assert.Equal(t, expected, res)
}

func Test_Sum_2(t *testing.T) {
	testTime := NewLocalTime(1, 1, 1)
	value := NewLocalTime(1, 121, 120)
	expected := NewLocalTime(4, 4, 1)

	res := testTime.Sum(value)
	assert.Equal(t, expected, res)
}

func Test_ToString(t *testing.T) {
	testTime := NewLocalTime(1, 1, 1)
	expected := "01:01:01"

	res := testTime.ToString()

	assert.Equal(t, expected, res)
}
