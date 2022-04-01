package demo

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func BenchmarkSHA1(b *testing.B) {
	data := []byte("Mary had a little lamb")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
}

func BenchmarkSHA256(b *testing.B) {
	data := []byte("Mary had a little lamb")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
}

func BenchmarkSHA512(b *testing.B) {
	data := []byte("Mary had a little lamb")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sha512.Sum512(data)
	}
}

func BenchmarkSHA512_Parallel(b *testing.B) {
	data := []byte("Mary had a little lamb")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sha512.Sum512(data)
		}
	})
}

func BenchmarkSHA512_withAlloc(b *testing.B) {
	data := []byte("Mary had a little lamb")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := sha512.New()
		h.Sum(data)
	}
}
