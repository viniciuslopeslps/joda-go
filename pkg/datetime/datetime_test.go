package datetime

import (
	"testing"

	"joda-go/pkg/date"
	"joda-go/pkg/localtime"

	"github.com/stretchr/testify/assert"
)

func Test_Time(t *testing.T) {
	expected := "20/01/2022 01:02:50"

	datetime := NewDateTime(20, 1, 2022, 1, 2, 50).ToString()

	assert.Equal(t, expected, datetime)
}

func Test_CropTime(t *testing.T) {
	expected := date.NewDate(20, 1, 2022)

	datetime := NewDateTime(20, 1, 2022, 1, 2, 50).CropTime()

	assert.Equal(t, expected, datetime)
}

func Test_CropDate(t *testing.T) {
	expected := localtime.NewLocalTime(1, 2, 50)

	datetime := NewDateTime(20, 1, 2022, 1, 2, 50).CropDate()

	assert.Equal(t, expected, datetime)
}
