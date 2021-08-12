package main

import (
	"fmt"
	"os"
)

// === Exercises === //

// ---------------------------------------------------------
// EXERCISE: Print Your Age
//  Print your age using Prinft
// EXPECTED OUTPUT
//  I'm 30 years old.
// NOTE
//  You should change 30 to your age, of course.
// ---------------------------------------------------------

// func main() {
// 	age := 28
// 	fmt.Printf("I'm %d years old.\n", age)
// }

// ---------------------------------------------------------
// EXERCISE: Print Your Name and LastName
//  Print your name and lastname using Printf
// EXPECTED OUTPUT
//  My name is Inanc and my lastname is Gumus.
// BONUS
//  Store the formatting specifier (first argument of Printf)
//    in a variable.
//  Then pass it to printf
// ---------------------------------------------------------

// func main() {
// 	first, last := "Tobiathus", "Eriktus"
// 	// BONUS: Use a variable for the format specifier
// 	f := "%s %s\n"

// 	fmt.Printf(f, first, last)
// }

// ---------------------------------------------------------
// EXERCISE: False Claims
//
//  Use printf to print the expected output using a variable.
//
// EXPECTED OUTPUT
//  These are false claims.
// ---------------------------------------------------------

// func main() {
// 	tf := false

// 	// TYPE YOUR CODE HERE
// 	fmt.Printf("These are %t claims.\n", tf)
// }

// ---------------------------------------------------------
// EXERCISE: Print the Temperature
//
//  Print the current temperature in your area using Printf
//
// NOTE
//  Do not use %v verb
//  Output "shouldn't" be like 29.500000
//
// EXPECTED OUTPUT
//  Temperature is 29.5 degrees.
// ---------------------------------------------------------

// func main() {
// 	temp := 64.4
// 	fmt.Printf("Temperature is %.2f degrees.\n", temp)
// }

// ---------------------------------------------------------
// EXERCISE: Double Quotes
//
//  Print "hello world" with double-quotes using Printf
//  (As you see here)
//
// NOTE
//  Output "shouldn't" be just: hello world
//
// EXPECTED OUTPUT
//  "hello world"
// ---------------------------------------------------------

// func main() {
// 	hw := "Hello world"
// 	fmt.Printf("%q", hw)
// }

// ---------------------------------------------------------
// EXERCISE: Print the Type
//
//  Print the type and value of 3 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3 is int
// ---------------------------------------------------------

// func main() {
// 	num := 3
// 	fmt.Printf("Type of %d is %[1]T\n", num)
// }

// ---------------------------------------------------------
// EXERCISE: Print the Type #2
//
//  Print the type and value of 3.14 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3.14 is float64
// ---------------------------------------------------------

// func main() {
// 	pi := 3.14
// 	fmt.Printf("Type of %.2f is %[1]T\n", pi)
// }

// ---------------------------------------------------------
// EXERCISE: Print the Type #3
//
//  Print the type and value of "hello" using fmt.Printf
//
// EXPECTED OUTPUT
// 	Type of hello is string
// ---------------------------------------------------------

// func main() {
// 	h := "hello"
// 	fmt.Printf("Type of %s is %[1]T.\n", h)
// }

// ---------------------------------------------------------
// EXERCISE: Print the Type #4
//  Print the type and value of true using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of true is bool
// ---------------------------------------------------------

// func main() {
// 	b := true
// 	fmt.Printf("Type of %t is %[1]T.\n", b)
// }

// ---------------------------------------------------------
// EXERCISE: Print Your Fullname
//
//  1. Get your name and lastname from the command-line
//  2. Print them using Printf
//
// EXAMPLE INPUT
//  Inanc Gumus
//
// EXPECTED OUTPUT
//  Your name is Inanc and your lastname is Gumus.
// ---------------------------------------------------------

func main() {
	first := os.Args[1]
	last := os.Args[2]
	f := "Your name is %s and your last name is %s.\n"
	fmt.Printf(f, first, last)
}
