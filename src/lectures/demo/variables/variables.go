package main

import "fmt"

func main() {
	var myName = "Mbiplang"
	fmt.Println("My name is", myName)

	var name string = "Kathy"
	fmt.Println("name =", name)

	// create and assign
	userName := "admin"
	fmt.Println("Username =", userName)

	var sum int
	fmt.Println("The sum is", sum)

	// compound assignment
	part1, other := 1, 5
	fmt.Println("Part 1 is", part1, "other is", other)

	// ok arrow idiom
	part2, other := 2, 0
	fmt.Println("Part 2 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("Sum is", sum)

	// Block assignment
	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("Lesson Name =", lessonName)
	fmt.Println("Lesson Type =", lessonType)

	// Compound assignment where we ignore one of the variables (using create and assign)
	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1, word2)
}
