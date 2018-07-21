package main

import (
	"fmt"
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
	people := []Person{
		Person{
			FirstName: "Nuttapon",
			LastName:  "Yodkaew",
			Age:       33,
		},
		Person{
			FirstName: "Atcha",
			LastName:  "Yodkaew",
			Age:       32,
		},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
}
