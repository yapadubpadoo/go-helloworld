package main

import (
	"fmt"

	"github.com/yapadubpadoo/gotools"
	"github.com/yapadubpadoo/gotools/stringhelper"
)

func init() {
	fmt.Println("Init!")
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(gotools.Add(10, 2))
	fmt.Println(stringhelper.Upper("cat"))
	fmt.Println(stringhelper.Concat("cat", "dog"))
}
