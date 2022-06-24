package localtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sum(t *testing.T) {
	testTime := NewLocalTime(1, 1, 1)
	value := NewLocalTime(1, 1, 1)
	expected := NewLocalTime(2, 2, 2)

	res := testTime.Sum(value)
	assert.Equal(t, expected, res)
}
