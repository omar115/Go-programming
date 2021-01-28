package main

import "fmt"

func main() {
	i := 1
	for i <= 3 { //basic type loop
		fmt.Println(i)
		i = i + 1
	}
	for j := 7; j <= 9; j++ { //generic for loop
		fmt.Println(j)
	}
	for { //this will run infinite loop until we do not break
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ { //using continue statement in loop
		if n%2 == 0 {
			continue

		}
		fmt.Println(n)
	}
}
