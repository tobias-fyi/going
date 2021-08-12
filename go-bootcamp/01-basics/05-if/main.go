package main

import (
	"fmt"
	"os"
)

// === Exercises === //

// ---------------------------------------------------------
// EXERCISE: Simplify It
//
//  Can you simplify the if statement inside the code below?
//
//  When:
//    isSphere == true and
//    radius is equal or greater than 200
//
//    It will print "It's a big sphere."
//
//    Otherwise, it will print "I don't know."
//
// EXPECTED OUTPUT
//  It's a big sphere.
// ---------------------------------------------------------

// func main() {
// 	// DO NOT TOUCH THIS
// 	isSphere, radius := true, 200

// 	var big bool

// 	if radius >= 200 {
// 		big = true
// 	}

// 	if big {
// 		fmt.Println("I don't know.")
// 	} else if isSphere {
// 		fmt.Println("It's a big sphere.")
// 	} else {
// 		fmt.Println("I don't know.")
// 	}
// }

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  You can also use strings.ContainsAny. Check out: https://golang.org/pkg/strings/#ContainsAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

// func main() {
// 	vowels := "aeiou"
// 	semi := "yw"

// 	var letter string

// 	if len(os.Args) < 2 {
// 		fmt.Println("Give me a letter.")
// 	} else {
// 		letter = os.Args[1]

// 		if len(letter) == 0 || len(letter) > 1 {
// 			fmt.Println("Give me a letter.")
// 		} else if strings.ContainsAny(letter, vowels) {
// 			fmt.Printf("'%s' is a vowel.\n", letter)
// 		} else if strings.ContainsAny(letter, semi) {
// 			fmt.Printf("'%s' is sometimes a vowel, sometimes not.\n", letter)
// 		} else {
// 			fmt.Printf("'%s' is a consonant.\n", letter)
// 		}
// 	}

// }

// ---------------------------------------------------------
// Challenge: LoginToMe
// - takes username and password, checking them against a
//   single user record in the database

// const (
// 	// Existing credentials
// 	user     = "tobes"
// 	pass     = "mcgobes"
// 	usage    = "LoginToMe Usage: [username] [password]\n\n - Enter user/pass\n - Profit!\n\nK then. Bye!\n"
// 	errUser  = "%s - Access denied.\n"
// 	errPass1 = "Please provide a password.\n"
// 	errPass2 = "Incorrect password for %s.\n"
// 	grant    = "%s - Access granted.\n"
// )

// func main() {
// 	// If nothing is passed, display usage info and quit
// 	if len(os.Args) < 2 {
// 		fmt.Printf("%s", usage)
// 	} else if len(os.Args) == 2 { // If only username provided
// 		fmt.Println(errPass1)
// 	} else if len(os.Args) > 2 { // If both provided
// 		// Save username and password provided from command line
// 		u, p := os.Args[1], os.Args[2]
// 		// Check username and prompt accordingly
// 		if u == user {
// 			if p == pass {
// 				fmt.Printf(grant, u)
// 			} else {
// 				// Check password and prompt accordingly
// 				fmt.Printf(errPass2, u)
// 			}
// 		} else {
// 			fmt.Printf(errUser, u)
// 		}
// 	}
// }

// ---------------------------------------------------------
// Challenge: LoginToMe #2
// - takes username and password, checking them against
//   multiple user records in the database.

const (
	// Existing credentials
	user1 = "tobes"
	pass1 = "mcgobes"
	user2 = "joe"
	pass2 = "virt"
	// Messages
	usage    = "LoginToMe Usage: [username] [password]\n\n - Enter user/pass\n - Profit!\n\nK then. Bye!\n"
	errUser  = "%s - Access denied.\n"
	errPass1 = "Please provide a password.\n"
	errPass2 = "Incorrect password for %s.\n"
	grant    = "%s - Access granted.\n"
)

// Insert usernames and passwords into map
var accounts = map[string]string{user1: pass1, user2: pass2}

func main() {
	if len(os.Args) < 2 { // If nothing is passed, display usage info and quit
		fmt.Printf("%s", usage)
	} else if len(os.Args) == 2 { // If only username provided
		fmt.Println(errPass1)
	} else if len(os.Args) > 2 { // If both provided
		// Save username and password provided from command line
		u, p := os.Args[1], os.Args[2]

		// Check if user exists in db
		if pass, ok := accounts[u]; ok {
			if pass == p {
				fmt.Printf(grant, u)
			} else {
				fmt.Printf(errPass2, u)
			}
		} else {
			fmt.Printf(errUser, u)
		}
	}
}
