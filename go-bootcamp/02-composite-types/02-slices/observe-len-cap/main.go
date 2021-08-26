package main

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Observe the length and capacity
//
//  Follow the instructions inside the code below to
//  gain more intuition about the length and capacity of a slice.
//
// ---------------------------------------------------------

func main() {
	// // --- #1 ---
	// // 1. create a new slice named: games
	// // var games []string

	// // 2. print the length and capacity of the games slice
	// // fmt.Printf("games - len: %d; cap: %d\n", len(games), cap(games))

	// // 3. comment out the games slice then declare it as an empty slice
	// // 4. print the length and capacity of the games slice
	// games := []string{}
	// fmt.Printf("games - len: %d; cap: %d\n", len(games), cap(games))

	// // 5. append the elements: "pacman", "mario", "tetris", "doom"
	// games = append(games, "pacman", "mario", "tetris", "doom")

	// // 6. print the length and capacity of the games slice
	// fmt.Printf("games - len: %d; cap: %d\n", len(games), cap(games))
	// // 7. comment out everything

	// 8. declare the games slice again using a slice literal
	//    (use the same elements from step 3)
	games := []string{"pacman", "mario", "tetris", "doom"}

	// --- #2 ---
	// 1. use a loop from 0 to 4 to slice the games slice, element by element.
	// 2. print its length and capacity along the way (in the loop).

	// for i := 0; i < 4; i++ {
	// 	end := 4 - i
	// 	games = games[:end]
	// 	fmt.Printf("games[:%d]'s len: %d cap: %d\n", end, len(games), cap(games))
	// }

	for i := 0; i < 4; i++ {
		games = games[:i]
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(games), cap(games))
	}

	fmt.Println()
	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	zero := games[:0]
	// 2. print the games and the new slice's len and cap
	fmt.Printf("games - len: %d; cap: %d\n", len(games), cap(games))
	fmt.Printf("zero - len: %d; cap: %d\n", len(zero), cap(zero))

	// 3. append a new element to the new slice
	zero = append(zero, "diablo 2")
	// 4. print the new slice's lens and caps
	fmt.Printf("zero - len: %d; cap: %d\n", len(zero), cap(zero))
	// 5. repeat the last two steps 5 times (use a loop)
	// 6. notice the growth of the capacity after the 5th append

	// Using a simple number-based loop
	// for i := 0; i < 5; i++ {
	// 	worms := "worms " + strconv.Itoa(i)
	// 	zero = append(zero, worms)
	// 	fmt.Printf("zero - len: %d; cap: %d | ", len(zero), cap(zero))
	// 	// fmt.Printf("%v\n", zero)
	// 	fmt.Println(zero) // Prints the same as above
	// }

	// Use this slice's elements to append to the new slice:
	more := []string{"ultima", "dagger", "pong", "coldspot", "zetra"}

	for i := range more {
		zero = append(zero, more[i])
		fmt.Printf("zero - len: %d; cap: %d | ", len(zero), cap(zero))
		fmt.Printf("%v\n", zero)
	}

	fmt.Println()

	// --- #4 ---
	// using a range loop, slice the zero slice element by element,
	// and print its length and capacity along the way.
	// observe that, the range loop only loops for the length, not the cap.
	// zLen := len(zero)
	// for i := 0; i < zLen; i++ {
	// 	zero = zero[1:]
	// 	fmt.Printf("zero - len: %d; cap: %d | ", len(zero), cap(zero))
	// 	fmt.Println(zero)
	// }

	zLen := len(zero)
	for i := 0; i < zLen; i++ {
		zero = zero[:i]
		fmt.Printf("zero - len: %d; cap: %d | ", len(zero), cap(zero))
		fmt.Println(zero)
	}

	// --- #5 ---
	// 1. do the 3rd step above again but this time, start by slicing
	//    the zero slice up to its capacity (use the cap function).
	// 2. print the elements of the zero slice in the loop.
	fmt.Println()

	zero = zero[:cap(zero)]
	zLen = len(zero)
	for i := 0; i < zLen; i++ {
		zero = zero[:i]
		fmt.Printf("zero - len: %d; cap: %d | ", len(zero), cap(zero))
		fmt.Println(zero)
	}

	// --- #6 ---
	// 1. change the one of the elements of the zero slice
	// 2. change the same element of the games slice
	// 3. print the games and the zero slices
	// 4. observe that they don't have the same backing array
	fmt.Println()

	zero[2] = "worms armageddon"
	games[2] = "worms 2"

	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)

	// --- #7 ---
	// try to slice the games slice beyond its capacity
	// no := games[:cap(games)+1]
	// fmt.Printf("%q\n", no)
	// panic: runtime error: slice bounds out of range [:5] with capacity 4
}
