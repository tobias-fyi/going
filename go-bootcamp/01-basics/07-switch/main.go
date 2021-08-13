package main

import (
	"fmt"
	"os"
)

// === Exercises === //

// ---------------------------------------------------------
// STORY
//  You're curious about the richter scales. When reporters
//  say: "There's been a 5.5 magnitude earthquake",
//  You want to know what that means.
//  So, you've decided to write a program to do that for you.
//
// EXERCISE: Richter Scale
//
//  1. Get the earthquake magnitude from the command-line.
//  2. Display its corresponding description.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 and less than 3.0        very minor
// 3.0 and less than 4.0        minor
// 4.0 and less than 5.0        light
// 5.0 and less than 6.0        moderate
// 6.0 and less than 7.0        strong
// 7.0 and less than 8.0        major
// 8.0 and less than 10.0       great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me the magnitude of the earthquake.
//
//  go run main.go ABC
//    I couldn't get that, sorry.
//
//  go run main.go 0.5
//    0.50 is micro
// ...
// ---------------------------------------------------------

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Kindly provide magnitude.")
// 		return
// 	}

// 	m, err := strconv.ParseFloat(os.Args[1], 64)

// 	if err != nil {
// 		fmt.Println("Invalid magnitude.")
// 		return
// 	}

// 	var msg string

// 	switch true {
// 	case m < 2.0:
// 		msg = "micro"
// 	case m < 3.0:
// 		msg = "very minor"
// 	case m < 4.0:
// 		msg = "minor"
// 	case m < 5.0:
// 		msg = "light"
// 	case m < 6.0:
// 		msg = "moderate"
// 	case m < 7.0:
// 		msg = "strong"
// 	case m < 8.0:
// 		msg = "major"
// 	case m < 10.0:
// 		msg = "great"
// 	default:
// 		msg = "massive"
// 	}

// 	fmt.Printf("%.2f is %s.\n", m, msg)
// }

// ---------------------------------------------------------
// EXERCISE: Convert
//
//  Convert the if statement to a switch statement.
// ---------------------------------------------------------

const (
	usage       = "Usage: [username] [password]"
	errUser     = "Access denied for %q.\n"
	errPwd      = "Invalid password for %q.\n"
	accessOK    = "Access granted to %q.\n"
	user, user2 = "jack", "sparrow"
	pass, pass2 = "1888", "1879"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	// REFACTOR THIS TO A SWITCH
	// if u != user && u != user2 {
	// 	fmt.Printf(errUser, u)
	// } else if u == user && p == pass {
	// 	fmt.Printf(accessOK, u)
	// } else if u == user2 && p == pass2 {
	// 	fmt.Printf(accessOK, u)
	// } else {
	// 	fmt.Printf(errPwd, u)
	// }

	switch {
	case u != user && u != user2:
		fmt.Printf(errUser, u)
	case u == user && p == pass:
		// fmt.Printf(accessOK, u)
		fallthrough // No need to duplicate above line
	case u == user2 && p == pass2:
		fmt.Printf(accessOK, u)
	default:
		fmt.Printf(errPwd, u)
	}
}
