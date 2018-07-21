package main

import (
	"fmt"
	"reflect"
)

func init() {
	fmt.Println("Init!")
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	me := Person{
		FirstName: "Nuttapon",
		LastName:  "Yodkaew",
		Age:       33,
	}
	fmt.Printf("%+v %s\n", me, reflect.TypeOf(me))
}
