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
	var number int
	for i := 1; i <= 100; i++ {
		number = i
		switch i := number; {
		//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
		case (i%3 == 0) && (i%5 == 0):
			fmt.Println("FizzBuzz")
			//  - Print "Fizz" if the integer is divisible by 3
		case i%3 == 0:
			fmt.Println("Fizz")
			//  - Print "Buzz" if the integer is divisible by 5
		case i%5 == 0:
			fmt.Println("Buzz")
			//* Print integers 1 to 50, except:
		default:
			fmt.Println(i)
		}
	}
}
