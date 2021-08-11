package main

import (
	"fmt"
	"os"
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

// func main() {
// 	var (
// 		planet string
// 		isTrue bool
// 		temp   float64
// 	)

// 	planet, isTrue, temp = "Mars", true, 19.5

// 	fmt.Println("Air is good on", planet)
// 	fmt.Println("It's", strconv.FormatBool(isTrue))
// 	fmt.Println("It's", temp, "degrees")
// 	// fmt.Printf("%v\n", temp)
// }

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
// 	_, b := multi()
// 	fmt.Println(b)
// }

// // multi is a function that returns multiple int values
// func multi() (int, int) {
// 	return 5, 4
// }

// ---------------------------------------------------------
// EXERCISE: Swapper
//  1. Change `color` to "orange" and `color2` to "green" at the same time
//     (use multiple-assignment)
//  2. Print the variables
// EXPECTED OUTPUT
//  orange green
// ---------------------------------------------------------

// func main() {
// 	color, color2 := "red", "blue"
// 	color, color2 = "orange", "green"
// 	fmt.Println(color, color2)
// }

// ---------------------------------------------------------
// EXERCISE: Swapper #2
//  1. Swap the values of `red` and `blue` variables
//  2. Print them
// EXPECTED OUTPUT
//  blue red
// ---------------------------------------------------------

// func main() {
// 	red, blue := "red", "blue"
// 	blue, red = red, blue
// 	fmt.Println(red, blue)
// }

// ---------------------------------------------------------
// EXERCISE: Discard The File
//  1. Print only the directory using `path.Split`
//  2. Discard the file part
// RESTRICTION
//  Use short declaration
// EXPECTED OUTPUT
//  secret/
// ---------------------------------------------------------

// func main() {
// 	dir, _ := path.Split("secret/file.txt")
// 	fmt.Println(dir)
// }

// ---------------------------------------------------------
// EXERCISE: Convert and Fix
//  Fix the code by using the conversion expression.
// EXPECTED OUTPUT
//  15.5
// ---------------------------------------------------------

// func main() {
// 	a, b := 10, 5.5
// 	fmt.Println(float64(a) + b)
// }

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #2
//  Fix the code by using the conversion expression.
// EXPECTED OUTPUT
//  10.5
// ---------------------------------------------------------

// func main() {
// 	a, b := 10, 5.5
// 	a = int(b)
// 	fmt.Println(float64(a) + b)
// }

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #3
//  Fix the code.
// EXPECTED OUTPUT
//  5.5
// ---------------------------------------------------------

// func main() {
// 	fmt.Println(5.5)
// }

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #4
//  Fix the code.
// EXPECTED OUTPUT
//  9.5
// ---------------------------------------------------------

// func main() {
// 	age := 2
// 	fmt.Println(7.5 + float64(age))
// }

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #5
//  Fix the code.
// HINTS
//   maximum of int8  can be 127
//   maximum of int16 can be 32767
// EXPECTED OUTPUT
//  1127
// ---------------------------------------------------------

// func main() {
// 	min := int8(127)
// 	max := int16(1000)
// 	fmt.Println(int16(min) + max)
// }

// ---------------------------------------------------------
// EXERCISE: Count the Arguments
//  Print the count of the command-line arguments
// INPUT
//  bilbo balbo bungo
// EXPECTED OUTPUT
//  There are 3 names.
// ---------------------------------------------------------

// func main() {
// 	count := len(os.Args) - 1
// 	fmt.Printf("There are %d names.\n", count)
// }

// ---------------------------------------------------------
// EXERCISE: Print the Path
//  Print the path of the running program by getting it from `os.Args` variable.
// HINT
//  Use `go build` to build your program.
//  Then run it using the compiled executable program file.
// EXPECTED OUTPUT SHOULD INCLUDE THIS
//  myprogram
// ---------------------------------------------------------

// func main() {
// 	fmt.Println(os.Args[0])
// }

// ---------------------------------------------------------
// EXERCISE: Print Your Name
//  Get it from the first command-line argument
// INPUT
//  Call the program using your name
// EXPECTED OUTPUT
//  It should print your name
// EXAMPLE
//  go run main.go inanc
//    inanc
// BONUS: Make the output like this:
//  go run main.go inanc
//    Hi inanc
//    How are you?
// ---------------------------------------------------------

// func main() {
// 	name := os.Args[1]
// 	fmt.Println("Hello,", name, "\nHow are you?")
// }

// ---------------------------------------------------------
// EXERCISE: Greet More People
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
// INPUT
//  bilbo balbo bungo
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

// func main() {
// 	count := len(os.Args) - 1
// 	fmt.Println("There are", count, "people!")
// 	for i := 1; i <= count; i++ {
// 		fmt.Println("Hello, great", os.Args[i])
// 	}
// 	fmt.Println("Nice to meet you all.")
// }

// ---------------------------------------------------------
// EXERCISE: Greet 5 People
//  Greet 5 people this time.
//  Please do not copy paste from the previous exercise!
// RESTRICTION
//  This time do not use variables.
// INPUT
//  bilbo balbo bungo gandalf legolas
// EXPECTED OUTPUT
//  There are 5 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Hello great gandalf !
//  Hello great legolas !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	fmt.Println("There are", len(os.Args)-1, "people!")
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("Hello, great", os.Args[i])
	}
	fmt.Println("Mice to neet you all.")
}
