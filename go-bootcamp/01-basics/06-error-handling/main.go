package main

import (
	"fmt"
	"time"
)

// === Notes === //

// func main() {
// 	var err error
// 	// Cannot use `if err {...}` syntax
// 	if err != nil {
// 		fmt.Printf("%q\n", err)
// 	}
// }

// === Exercises === //

// -------------------------------
// Feet to Meters
//

// func main() {
// 	conv := 0.3048 // Feet -> Meters

// 	// Handle case when nothing is passed
// 	if len(os.Args) < 2 {
// 		fmt.Println("Please enter a valid value to be converted.")
// 		return
// 	}

// 	arg := os.Args[1]
// 	feet, err := strconv.ParseFloat(arg, 64)
// 	if err != nil {
// 		fmt.Printf("Error: %q is not a number.", arg)
// 		return
// 	}
// 	meters := float64(feet) * conv
// 	fmt.Printf("%g feet is %g meters.\n", feet, meters)
// }

// ---------------------------------------------------------
// STORY
//
//  Your boss wants you to create a program that will check
//  whether a person can watch a particular movie or not.
//
//  Imagine that another program provides the age to your
//  program. Depending on what you return, the other program
//  will issue the tickets to the person automatically.
//
// EXERCISE: Movie Ratings
//
//  1. Get the age from the command-line.
//
//  2. Return one of the following messages if the age is:
//     -> Above 17         : "R-Rated"
//     -> Between 13 and 17: "PG-13"
//     -> Below 13         : "PG-Rated"
//
// RESTRICTIONS:
//  1. If age data is wrong or absent let the user know.
//  2. Do not accept negative age.
//
// BONUS:
//  1. BONUS: Use if statements only twice throughout your program.
//  2. BONUS: Use an if statement only once.
//
// EXPECTED OUTPUT
//  go run main.go 18
//    R-Rated
//
//  go run main.go 17
//    PG-13
//
//  go run main.go 12
//    PG-Rated
//
//  go run main.go
//    Requires age
//
//  go run main.go -5
//    Wrong age: "-5"
// ---------------------------------------------------------

// func main() {
// 	if args := os.Args; len(args) != 2 {
// 		fmt.Println("Please provide age")

// 	} else if age, err := strconv.Atoi(args[1]); err != nil || age < 0 {
// 		fmt.Println("Invalid age:", args[1])

// 	} else if age < 13 {
// 		fmt.Println("PG")

// 	} else if age < 17 {
// 		fmt.Println("PG-13")

// 	} else {
// 		fmt.Println("R")

// 	}
// }

// ---------------------------------------------------------
// Challenge: Parts of a Day
//

func main() {
	// switch h := time.Now().Hour(); true {
	// case h > 5 && h < 12:
	// 	fmt.Println("Good morning!")
	// case h > 12 && h < 17:
	// 	fmt.Println("Good afternoon!")
	// case h > 17 && h < 20:
	// 	fmt.Println("Good evening!")
	// case h > 20 || h < 5:
	// 	fmt.Println("Good night!")
	// }

	// Better
	switch h := time.Now().Hour(); true {
	case h >= 18:
		fmt.Println("Good evening!")
	case h >= 12:
		fmt.Println("Good afternoon!")
	case h >= 5:
		fmt.Println("Good morning!")
	default:
		fmt.Println("Good night!")
	}

}
