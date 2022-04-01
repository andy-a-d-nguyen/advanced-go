package demo_test

import (
	"demo"
	"testing"
)

func TestAdd(t *testing.T) {
	l, r := 1, 2
	expect := 3
	got := demo.Add(1, 2)
	if got != expect {
		t.Errorf("Expected %v when adding %v and %v. Got %v", expect, l, r, got)
	}
}
