package main

import (
	"fmt"
	"reflect"
)

func init() {
	fmt.Println("Init!")
}

func main() {
	const (
		myFirstname = "Nuttapon"
		myLastname  = "Yodkaew"
	)
	a := "hello"
	b := 10
	c := 10.109999999999999999999999999999999
	d := true
	fmt.Printf("%s %d %f %v\n", a, b, c, d)

	fmt.Printf("%s is %s\n", a, reflect.TypeOf(a))

}
