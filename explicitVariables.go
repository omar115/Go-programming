/* here we will see how to declare
variable explicitly and in different ways
*/

package main

import "fmt"

func main() {
	var a = "initial" //var declares one or more variables
	fmt.Println(a)

	var b, c int = 1, 2 //can declare mutliple variables at same time
	fmt.Println(b, c)

	var d = true // will infer the type of variable
	fmt.Println(d)

	var e int // variable decalred without initialization are zero valued
	fmt.Println(e)

	f := "apple" //:= is a shortcut way to declare and initialize the variable
	fmt.Println(f)
}
