//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func sayHello(name string) {
	fmt.Println("Hello,", name, "nice to meet you!")
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func crazyFunc() string {
	return "This is a message to all developers, learn golang!!!"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func add3Numbers(x, y, z int) int {
	return x + y + z
}

//* Write a function that returns any number
func returnAnyNumber() int {
	return 30
}

//* Write a function that returns any two numbers
func returnAny2Numbers() (int, int) {
	return 30, 100
}

//* Add three numbers together using any combination of the existing functions.
// func add3Numbers(a, b, c int) int {
// 	return a + b + c
// }

//  * Print the result
//* Call every function at least once

func main() {

	sayHello("Mbiplang")
	fmt.Println(crazyFunc())
	fmt.Println("Sum of the 3 numbers is:", add3Numbers(43, 23, 5))
	fmt.Println("Any number:", returnAnyNumber())
	fmt.Println(returnAny2Numbers())
	first, second := returnAny2Numbers()
	fmt.Println(add3Numbers(first, second, returnAnyNumber()))
}
