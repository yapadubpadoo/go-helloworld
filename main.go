package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init!")
}

func main() {
	defer fmt.Println("World")
	defer fmt.Println("Middle")
	fmt.Println("Hello")
}
