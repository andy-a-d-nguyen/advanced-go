- case for generics
    - how do we create generic behavior today?
    - summation function that uses type switch

- build adder function, highlighting need for checks and guards
- create driver in main and run to show that it works

```go
package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	data := [][]interface{}{
		[]interface{}{1, 2, 3},
		[]interface{}{3.14, 6.02},
		[]interface{}{"foo", "bar", "baz"},
		nil,
		[]interface{}{1 + 2i, 2 + 3i},
	}
	for _, d := range data {
		sum, err := add(d)
		if err != nil {
			log.Println(fmt.Errorf("failed to add values '%v': %q", d, err))
			continue
		}
		fmt.Printf("Sum of %v: %v\n", d, sum)
	}
}

func add(s []interface{}) (interface{}, error) {
	if len(s) == 0 {
		return nil, errors.New("provided slice cannot be empty")
	}
	switch s[0].(type) {
	case int:
		var sum int
		for _, v := range s {
			if i, ok := v.(int); ok {
				sum += i
			} else {
				return nil, errors.New("all values must be of same type")
			}
		}
		return sum, nil
	case float64:
		var sum float64
		for _, v := range s {
			if f, ok := v.(float64); ok {
				sum += f
			} else {
				return nil, errors.New("all values must be of same type")
			}
		}
		return sum, nil
	case string:
		var sum string
		for _, v := range s {
			if s, ok := v.(string); ok {
				sum += s
			} else {
				return nil, errors.New("all values must be of same type")

			}
		}
		return sum, nil
	default:
		return nil, errors.New("unsupported type")
	}
}
```