package main

import "fmt"

func main() {
	fmt.Println(add(3.14, 3))
	fmt.Println(add(1, 2))
}

type reals interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func add[T reals, S reals](a T, b S) float64 {
	return float64(a) + float64(b)
}
