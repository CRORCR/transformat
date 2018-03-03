package calc_test

import (
	"testing"
	"transformat/day1/structdmeo"
)

func TestName(t *testing.T) {
	add := structdmeo.Add(5, 1)
	if add == 6 {
		t.Fatal("add is failed")
	}
	t.Logf("add is success")
}
