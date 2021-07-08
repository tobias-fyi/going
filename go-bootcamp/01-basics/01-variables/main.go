package main

import (
	"fmt"
	"strconv"
)

// === Notes === //

// func main() {
// 	_, file := path.Split("go-bootcamp/01-basics/01-path-separator.go")
// 	fmt.Println("file:", file)
// }

// func main() {
// 	speed, force := 100, 2.5
// 	momentum := float64(speed) * force
// 	fmt.Println(momentum)
// }

// func main() {
// 	fmt.Println("Path:", os.Args[0])
// 	fmt.Println("One:", os.Args[1])
// 	fmt.Println("Last:", os.Args[len(os.Args)-1])
// }

// === Exercises === //

// ---------------------------------------------------------
// EXERCISE: Make It Blue
//  1. Change `color` variable's value to "blue"
//  2. Print it
// EXPECTED OUTPUT
//  blue
// ---------------------------------------------------------

// func main() {
// 	color := "green"
// 	color = "blue"
// 	fmt.Println(color)
// }

// ---------------------------------------------------------
// EXERCISE: Variables To Variables
//  1. Change the value of `color` variable to "dark green"
//  2. Do not assign "dark green" to `color` directly
//     Instead, use the `color` variable again
//     while assigning to `color`
//  3. Print it
// RESTRICTIONS
//  WRONG ANSWER, DO NOT DO THIS:
//  `color = "dark green"`
// HINT
//  + operator can concatenate string values
//  Example: "h" + "e" + "y" returns "hey"
// EXPECTED OUTPUT
//  dark green
// ---------------------------------------------------------

// func main() {
// 	color := "green"
// 	color = "dark " + color
// 	fmt.Println(color)
// }

// ---------------------------------------------------------
// EXERCISE: Assign With Expressions
//  1. Multiply 3.14 with 2 and assign it to `n` variable
//  2. Print the `n` variable
// HINT
//  Example: 3 * 2 = 6
//  * is the product operator (it multiplies numbers)
// EXPECTED OUTPUT
//  6.28
// ---------------------------------------------------------

// func main() {
// 	// Declares a new float64 variable
// 	n := 0.
// 	n = 3.14 * 2.
// 	fmt.Println(n)
// }

// ---------------------------------------------------------
// EXERCISE: Find the Rectangle's Perimeter
//  1. Find the perimeter of a rectangle
//     Its width is  5
//     Its height is 6
//  2. Assign the result to the `perimeter` variable
//  3. Print the `perimeter` variable
// HINT
//  Formula = 2 times the width and height
// EXPECTED OUTPUT
//  22
// BONUS
//  Find more formulas here and calculate them in new programs
//  https://www.mathsisfun.com/area.html
// ---------------------------------------------------------

// func main() {
// 	var (
// 		perimeter     int
// 		width, height = 5, 6
// 	)

// 	perimeter = 2 * (width + height)

// 	fmt.Println(perimeter)
// }

// ---------------------------------------------------------
// EXERCISE: Multi Assign
//  1. Assign "go" to `lang` variable
//     and assign 2 to `version` variable
//     using a multiple assignment statement
//  2. Print the variables
// EXPECTED OUTPUT
//  go version 2
// ---------------------------------------------------------

// func main() {
// 	var (
// 		lang    string
// 		version int
// 	)

// 	lang, version = "go", 2

// 	fmt.Println(lang, "version", version)
// }

// ---------------------------------------------------------
// EXERCISE: Multi Assign #2
//  1. Assign the correct values to the variables
//     to match to the EXPECTED OUTPUT below
//  2. Print the variables
// HINT
//  Use multiple Println calls to print each sentence.
// EXPECTED OUTPUT
//  Air is good on Mars
//  It's true
//  It is 19.5 degrees
// ---------------------------------------------------------

func main() {
	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, isTrue, temp = "Mars", true, 19.5

	fmt.Println("Air is good on", planet)
	fmt.Println("It's", strconv.FormatBool(isTrue))
	fmt.Println("It's", temp, "degrees")
	// fmt.Printf("%v\n", temp)
}

// ---------------------------------------------------------
// EXERCISE: Multi Short Func
// 	1. Declare two variables using multiple short declaration syntax
//  2. Initialize the variables using `multi` function below
//  3. Discard the 1st variable's value in the declaration
//  4. Print only the 2nd variable
// NOTE
//  You should use `multi` function while initializing the variables
// EXPECTED OUTPUT
//  4
// ---------------------------------------------------------

// func main() {
// 	// ADD YOUR DECLARATIONS HERE

// 	fmt.Println(b)
// }

// // multi is a function that returns multiple int values
// func multi() (int, int) {
// 	return 5, 4
// }
