package main

import (
	"strings"
	"testing"
)

func toUpper(s string) string {
	var b strings.Builder
	for i := range s {
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
