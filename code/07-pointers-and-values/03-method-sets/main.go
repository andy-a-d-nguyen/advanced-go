package main

import "fmt"

func main() {
	var num smartNumber
	v := myNum(42)
	num = &v
	fmt.Println(num, num.isEven())
	num.addOne()
	fmt.Println(num, num.isEven())
}

type myNum int

func (num myNum) isEven() bool {
	return num%2 == 0
}

func (num *myNum) addOne() {
	result := *num + 1
	*num = result
}

type smartNumber interface {
	isEven() bool
	addOne()
}
