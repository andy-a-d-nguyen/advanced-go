- intro to generics
    - issue 43651 (https://github.com/golang/go/issues/43651)
    - proposal doc - https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md
    - basic syntax

- replicating `copy` with basic generics
    - note use of `any` as standin for `interface{}` in type constraint

```go
package main

import (
	"fmt"
)

func main() {
	testScores := []float64 {
		87.3,
		105,
		63.5,
		27,
	}
	c := clone[float64](testScores)

	fmt.Println(c, &testScores[0] == &c[0])
}

func clone[V any](s []V) []V {
	result := make([]V, len(s))
	for i, v := range s {
		result[i] = v
	}

	return result
}
```

- more useful example - generic function to shallow clone a map
    - note use of `comparable` constraint for K

```go
package main

import (
	"fmt"
)

func main() {
	testScores := map[string]float64 {	
		"Harry": 87.3,
		"Hermione": 105,
		"Ronald": 63.5,
		"Neville": 27,
	}
	c := clone[string, float64](testScores)

	fmt.Println(c)
}

func clone[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}

	return result
}

```

