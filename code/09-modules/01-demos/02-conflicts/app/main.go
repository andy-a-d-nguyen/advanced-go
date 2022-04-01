package main

import (
	"fmt"

	"github.com/vansimke/dep/v2"
)

func main() {
	u := dep.User{
		Username:  "byoda",
		Firstname: "baby",
		Lastname:  "yoda",
	}

	fmt.Println(u)
}
