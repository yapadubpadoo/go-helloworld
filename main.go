package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init!")
}

func main() {
	a := 10
	b := 20

	switch {
	case a > 10:
		fmt.Println("a > 10")
	case b > 10:
		fmt.Println("b > 10")
	default:
		fmt.Println("wtf")
	}
}
