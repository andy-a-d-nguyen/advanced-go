package demo

import (
	"fmt"
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

func TestAdd_Table(t *testing.T) {
	cases := []struct {
		l, r   int
		expect int
	}{
		{l: 1, r: 2, expect: 3},
		{l: 0, r: 2, expect: 2},
		{l: 2, r: 0, expect: 2},
		{l: 0, r: 0, expect: 0},
	}
	for _, data := range cases {
		t.Run(fmt.Sprintf("Adding %v and %v. Expect: %v\n", data.l, data.r, data.expect), func(t *testing.T) {
			got := Add(data.l, data.r)
			if got != data.expect {
				t.Fail()
			}
		})
	}
}
