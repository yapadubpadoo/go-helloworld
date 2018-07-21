package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init!")
}

func main() {
	names := map[string]int{
		"a": 20,
		"b": 30,
		"c": 40,
	}

	for k, v := range names {
		fmt.Println(k, "->", v)
	}
}
