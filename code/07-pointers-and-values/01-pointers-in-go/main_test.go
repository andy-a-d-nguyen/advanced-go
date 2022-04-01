package main

import "testing"

func BenchmarkPassByValue(b *testing.B) {
	var arr [1000]int
	for i := 0; i < b.N; i++ {
		ch := make(chan [1000]int, 1000)
		for j := 0; j < 1000; j++ {
			ch <- arr
		}
	}
}

func BenchmarkPassByPointer(b *testing.B) {
	var arr [1000]int
	for i := 0; i < b.N; i++ {
		ch := make(chan *[1000]int, 1000)
		for j := 0; j < 1000; j++ {
			ch <- &arr
		}
	}
}
