package main

import (
	"fmt"
)


var a, b, c int

func main() {
	// var message string 
	// message := 20.3 
	// variables initialization

	// var char rune = 'a' char 
	// uint - can be only positive number not negative.

	a, b, c = 1, 2, 3

	a, _, c = c, b, a

	fmt.Println(a, b, c)
	print()

}

func print() {
	fmt.Println(a, b, c)
}

/* const - константа(неизменное) var - переменная(изменяемая) */
