package time

import (
	"testing"
)

func TestTime_Sum(t *testing.T) {
	testTime := NewTime(1, 1, 1)
	value := NewTime(1, 1, 1)
	expected := NewTime(2, 2, 2)
	res := testTime.Sum(value)

	if expected != res {
		t.Fatal("Values are different: %v != %v", expected, res)
	}
}
