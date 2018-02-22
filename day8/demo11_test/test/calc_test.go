package test

import (
	"oldboy/day8/demo11_test"
	"testing"
)

func TestAdd(t *testing.T) {
	result := demo11_test.Add(1, 2)
	if result != 10 {
		t.Fatal("值不是10")
		return
	}
	t.Logf("success")
}

func TestSub(t *testing.T) {
	result := demo11_test.Sub(10, 1)
	if result != 9 {
		t.Fatal("值不是8")
	}
	t.Logf("success")
}
