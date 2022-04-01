package main

import "fmt"

func main() {
	num := myNum(42)
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
