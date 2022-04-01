- example that uses generic type with type sets
    - revisit generic adder using type set

- write this up and highlight lack of guard code
    - also return types don't need assertions!

- update Addable with bool type and try to run to show that Go does check usages

```go
package main

import (
	"fmt"
)

type Addable interface {
	int | float64 | string
}

func main() {
	a1 := []int{1, 2, 3}
	a2 := []float64{3.14, 6.02}
	a3 := []string{"foo", "bar", "baz"}

	fmt.Printf("Sum of %v: %v\n", a1, add(a1))
	fmt.Printf("Sum of %v: %v\n", a2, add(a2))
	fmt.Printf("Sum of %v: %v\n", a3, add(a3))
}

func add[V Addable](s []V) V {
	var result V
	for _, v := range s {
		result += v
	}
	return result
}

```