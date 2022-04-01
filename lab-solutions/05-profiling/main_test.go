package main

import (
	"encoding/csv"
	"os"
	"testing"
)

func Benchmark(b *testing.B) {
	f, err := os.Open(fileName)
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	data, err := r.ReadAll()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		rawRecordCh = make(chan []string)
		recordCh = make(chan Record)
		resultCh = make(chan Result)
		errorCh = make(chan error)
		b.StartTimer()
		execute(data)
	}
}
