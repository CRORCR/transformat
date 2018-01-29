package test

import (
	"testing"
	"transformat/day1/demo1calc"
)

func TestAdd(t *testing.T) {
	sum := demo1calc.Add(3, 2)
	if sum != 5 {
		t.Fatal("add is not right, sum:%v expected:5", sum)
	}
	t.Logf("add is ok")
}
