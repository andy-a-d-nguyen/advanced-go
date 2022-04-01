package demo

import (
	"testing"
)

func TestAdd(t *testing.T) {
	l, r := 1, 2
	expect := 3
	got := Add(1, 2)
	if got != expect {
		t.Errorf("Expected %v when adding %v and %v. Got %v", expect, l, r, got)
	}
}

func TestSub(t *testing.T) {
	l, r := 10, 6
	expect := 4
	got := sub(l, r)
	if got != expect {
		t.Errorf("Expected %v when subtracting %v from %v. Got %v", expect, r, l, got)
	}
}
