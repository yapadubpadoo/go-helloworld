package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init!")
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	a := [...]int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
		b = append(b, a[i])
	}
}
