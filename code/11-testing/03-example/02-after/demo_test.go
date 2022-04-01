package demo

import (
	"fmt"
)

func Example() {
	l, r := 10, 20
	fmt.Println(Add(l, r))

	// Output:
	// 30
}

func ExampleAdd() {
	l, r := 1, 2
	fmt.Println(Add(l, r))

	// Output:
	// 3
}
