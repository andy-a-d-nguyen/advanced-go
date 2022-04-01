package main

import (
	"fmt"
	"runtime"
)

func main() {
	m := make(map[int]int)
	printHeapUse()
	for i := 0; i < 1e6; i++ {
		m[i] = i
	}
	printHeapUse()
	for i := 0; i < 1e6; i++ {
		delete(m, i)
	}
	printHeapUse()
	runtime.GC()
	printHeapUse()
	fmt.Println(m)
}

func printHeapUse() {
	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Printf("Heap memory in use: %v\n", stats.HeapInuse)
}
