package main

import (
	"strconv"
	"strings"
	"testing"
)

// write initial function without 'if' guards and then add them
// as the fuzz tests find failures
// really we should just use strings.ToUpper!
func toUpper(s string) string {
	var b strings.Builder
	for i := range s {
		if _, err := strconv.Atoi(string(s[i])); err == nil {
			b.WriteByte(s[i])
			continue
		}
		if s[i] >= 'A' && s[i] <= 'Z' {
			b.WriteByte(s[i])
			continue
		}
		b.WriteByte(byte(s[i]) - 32)
	}
	return b.String()
}

func TestUpper(t *testing.T) {
	have := "hello"
	expect := "HELLO"
	got := toUpper(have)
	t.Log(got, expect)
	if expect != got {
		t.Fail()
	}
}

func FuzzFoo(f *testing.F) {
	f.Add("hello")
	f.Fuzz(func(t *testing.T, s string) {
		out := toUpper(s)
		if out != strings.ToUpper(s) {
			t.Fail()
		}
	})
}
