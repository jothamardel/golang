//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {

	for i := 1; i <= 50; i++ {
		divisibleBy3 := i%3 == 0
		divisibleBy5 := i%5 == 0

		switch i := i; {
		//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
		case divisibleBy3 && divisibleBy5:
			fmt.Println("FizzBuzz")
			//  - Print "Fizz" if the integer is divisible by 3
		case divisibleBy3:
			fmt.Println("Fizz")
			//  - Print "Buzz" if the integer is divisible by 5
		case divisibleBy5:
			fmt.Println("Buzz")
			//* Print integers 1 to 50, except:
		default:
			fmt.Println(i)
		}
	}
}
