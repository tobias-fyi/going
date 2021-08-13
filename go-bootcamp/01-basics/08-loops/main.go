package main

import (
	"fmt"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Dynamic Table
//
//  Get the size of the table from the command-line
//    Passing 5 should create a 5x5 table
//    Passing 10 for a 10x10 table
//
// HINT
//  There was a max constant in the original program.
//  That determines the size of the table.
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Give me the size of the table
//
//  go run main.go -5 (or ABC)
//    Wrong size
//
//  go run main.go 2
//    X    0    1    2
//    0    0    0    0
//    1    0    1    2
//    2    0    2    4
// ---------------------------------------------------------

// func main() {
// 	// Get size of table from command line
// 	if len(os.Args) < 2 {
// 		fmt.Println("Please provide a table size.")
// 		return
// 	}

// 	size, err := strconv.Atoi(os.Args[1])
// 	if err != nil || size < 0 {
// 		fmt.Println("Wrong size.")
// 	}

// 	fmt.Printf("%5s", "X")
// 	for i := 0; i <= size; i++ {
// 		fmt.Printf("%5d", i)
// 	}
// 	fmt.Println()

// 	for i := 0; i <= size; i++ {
// 		fmt.Printf("%5d", i)

// 		for j := 0; j <= size; j++ {
// 			fmt.Printf("%5d", i*j)
// 		}

// 		fmt.Println()
// 	}
// }

// ---------------------------------------------------------
// EXERCISE: Math Tables
//
//  Create division, addition and subtraction tables
//
//  1. Get the math operation and
//     the size of the table from the user
//
//  2. Print the table accordingly
//
//  3. Correctly handle the division by zero error
//
//
// BONUS #1
//
//  Use strings.IndexAny function to detect
//    the valid operations.
//
//  Search on Google for: golang pkg strings IndexAny
//
// BONUS #2
//
//  Add remainder operation as well (remainder table using %).
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Usage: [op=*/+-] [size]
//
//  go run main.go "*"
//    Size is missing
//    Usage: [op=*/+-] [size]
//
//  go run main.go "%" 4
//    Invalid operator.
//    Valid ops one of: */+-
//
//  go run main.go "*" 4
//    *    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    2    3    4
//    2    0    2    4    6    8
//    3    0    3    6    9   12
//    4    0    4    8   12   16
//
//  go run main.go "/" 4
//    /    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    0    0    0
//    2    0    2    1    0    0
//    3    0    3    1    1    0
//    4    0    4    2    1    1
//
//  go run main.go "+" 4
//    +    0    1    2    3    4
//    0    0    1    2    3    4
//    1    1    2    3    4    5
//    2    2    3    4    5    6
//    3    3    4    5    6    7
//    4    4    5    6    7    8
//
//  go run main.go "-" 4
//    -    0    1    2    3    4
//    0    0   -1   -2   -3   -4
//    1    1    0   -1   -2   -3
//    2    2    1    0   -1   -2
//    3    3    2    1    0   -1
//    4    4    3    2    1    0
//
// BONUS:
//
//  go run main.go "%" 4
//    %    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    0    1    1    1
//    2    0    0    0    2    2
//    3    0    0    1    0    3
//    4    0    0    0    1    0
//
// NOTES:
//
//   When running the program in Windows Git Bash, you might need
//   to escape the characters like this.
//
//   This happens because those characters have special meaning.
//
//   Division:
//     go run main.go // 4
//
//   Multiplication and others:
//   (this is also necessary for non-Windows bashes):
//
//     go run main.go "*" 4
// ---------------------------------------------------------

// func main() {
// 	// Get size of table from command line
// 	if len(os.Args) < 3 {
// 		fmt.Println("Usage: [op=*/+-] [size]")
// 		return
// 	}

// 	op := os.Args[1]
// 	opIndex := strings.IndexAny(op, "*/+-%")
// 	if opIndex == -1 {
// 		fmt.Printf("Invalid operation: %s\n", op)
// 	}

// 	size, err := strconv.Atoi(os.Args[2])
// 	if err != nil || size < 0 {
// 		fmt.Println("Wrong size.")
// 	}

// 	fmt.Printf("%5s", op)
// 	for i := 0; i <= size; i++ {
// 		fmt.Printf("%5d", i)
// 	}
// 	fmt.Println()

// 	for i := 0; i <= size; i++ {
// 		fmt.Printf("%5d", i)

// 		for j := 0; j <= size; j++ {
// 			var cell int // uninit value is 0

// 			switch op {
// 			case "%":
// 				if j != 0 { // Handle div by zero
// 					cell = i % j
// 				}

// 			case "/":
// 				if j != 0 { // Handle div by zero
// 					cell = i / j
// 				}

// 			case "+":
// 				cell = i + j

// 			case "-":
// 				cell = i - j

// 			default:
// 				cell = i * j
// 			}
// 			fmt.Printf("%5d", cell)
// 		}

// 		fmt.Println()
// 	}
// }

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

// const (
// 	comErr = "Please provide valid min and max."
// 	numErr = "Invalid min or max."
// )

// func main() {
// 	// Get min/max from command line and check them
// 	if len(os.Args) < 3 {
// 		fmt.Println(comErr)
// 	}

// 	min, errMin := strconv.Atoi(os.Args[1])
// 	max, errMax := strconv.Atoi(os.Args[2])
// 	if errMin != nil || errMax != nil {
// 		fmt.Println(numErr)
// 	}

// 	// Loop thru and sum even numbers between min and max
// 	var sum int
// 	for i := min; i <= max; i++ {

// 		if i%2 != 0 { // Odd
// 			continue
// 		}

// 		if i == min || i == min+1 {
// 			fmt.Printf("%d", i)
// 		} else {
// 			fmt.Printf(" + %d", i)
// 		}
// 		sum += i
// 	}

// 	fmt.Printf(" = %d\n", sum)
// }

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

// const (
// 	comErr = "Please provide valid min and max."
// 	numErr = "Invalid min or max."
// )

// func main() {
// 	// Get min/max from command line and check them
// 	if len(os.Args) < 3 {
// 		fmt.Println(comErr)
// 	}

// 	min, errMin := strconv.Atoi(os.Args[1])
// 	max, errMax := strconv.Atoi(os.Args[2])
// 	if errMin != nil || errMax != nil {
// 		fmt.Println(numErr)
// 	}

// 	// Loop thru and sum even numbers between min and max
// 	var sum int
// 	i := min
// 	for {

// 		if i%2 != 0 { // Odd
// 			i++
// 			continue
// 		}

// 		if i == min || i == min+1 {
// 			fmt.Printf("%d", i)
// 		} else {
// 			fmt.Printf(" + %d", i)
// 		}
// 		sum += i

// 		if i == max {
// 			break
// 		}

// 		i++
// 	}

// 	fmt.Printf(" = %d\n", sum)
// }

// ---------------------------------------------------------
// EXERCISE: Infinite Kill
//
//  1. Create an infinite loop
//  2. On each step of the loop print a random character.
//  3. And, sleep for 250 milliseconds.
//  4. Run the program and wait a couple of seconds
//     then kill it using CTRL+C keys
//
// RESTRICTIONS
//  1. Print one of those characters randomly: \ / | -
//  2. Before printing a character print also this
//     escape sequence: \r
//
//     Like this: "\r/", or this: "\r|", and so on...
//
//  3. NOTE : If you're using Go Playground, use "\f" instead of "\r"
//
// HINTS
//  1. Use `time.Sleep` to sleep.
//  2. You should pass a `time.Duration` value to it.
//  3. Check out the Go online documentation for the
//     `Millisecond` constant.
//  4. When printing, do not use a newline! Or a Println!
//     Use Print or Printf instead.
//
// NOTE
//  If this exercise is hard for you now, wait until the
//  lucky number lecture. Even then so, then just skip it.
//
//  You can return back to it afterwards.
//
// EXPECTED OUTPUT
//  - The program should display the following messages in any order.
//  - And, the first character (\, -, /, or |) should be randomly
//    displayed.
//  - \r or \f sequence will clear the line.
//
//  \ Please Wait. Processing....
//  - Please Wait. Processing....
//  / Please Wait. Processing....
//  | Please Wait. Processing....
// ---------------------------------------------------------

// func main() {
// 	// Convert chars to runes
// 	chars := `\/|-`
// 	runes := []rune(chars)

// 	for {
// 		// Create random int to index runes
// 		randIndex := rand.Intn(len(runes))
// 		randChar := string(runes[randIndex])
// 		fmt.Printf("\r%s Please wait. Processing...", randChar)
// 		time.Sleep(250 * time.Millisecond)
// 	}
// }

// Same thing, but going through the runes in order
func main() {
	chars := `\|/-`
	runes := []rune(chars)
	var i int

	for { // Infinite loop thru runes in order
		i = i % len(runes)
		randChar := string(runes[i])
		fmt.Printf("\r%s Please wait. Processing...", randChar)
		time.Sleep(100 * time.Millisecond)
		i++
	}
}
