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
	age       int
}

func (p *Person) info() string {
	return fmt.Sprintf("%s %s, %d", p.FirstName, p.LastName, p.age)
}

func (p *Person) growUp() {
	p.age = p.age + 1
}

func main() {
	people := []Person{
		Person{
			FirstName: "Nuttapon",
			LastName:  "Yodkaew",
			age:       33,
		},
		Person{
			FirstName: "Atcha",
			LastName:  "Yodkaew",
			age:       32,
		},
	}

	for _, person := range people {
		// result, _ := json.Marshal(person)
		// fmt.Printf("%v\n", string(result))
		person.growUp()
		fullName := person.info()
		fmt.Println(fullName)
	}
}
