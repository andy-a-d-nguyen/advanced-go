package main

import (
	"errors"
	"fmt"
)

type myError struct {
	msg          string
	wrappedError error
}

func (me myError) Error() string {
	return me.msg
}

func (me myError) Unwrap() error {
	return me.wrappedError
}

var errFoo = errors.New("Fixed errors allow us to check for 'known' errors")

func main() {
	err1 := errFoo
	err2 := myError{
		msg:          "Custom error types allow us to store additional information",
		wrappedError: err1,
	}
	err3 := fmt.Errorf("This error wraps err2: %w", err2)

	fmt.Printf("The top level error: %v\n", err3)
	fmt.Printf("Unwrap one level: %v\n", errors.Unwrap(err3))
	fmt.Printf("Is err3 wrapping fixedError? %v\n", errors.Is(err3, errFoo))

	var me myError
	ok := errors.As(err3, &me)
	fmt.Printf("\nRetrieve myError object from err3:\n%v\n%v\n", ok, me)

}
