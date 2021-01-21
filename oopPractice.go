package main

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

func (a author) details() {
	fmt.Println("Name is: ", a.fullName())
}

func main() {
	author1 := author{
		"Md Omar",
		"Hasan",
		"I love Python and Go",
	}

	author1.details()
}
