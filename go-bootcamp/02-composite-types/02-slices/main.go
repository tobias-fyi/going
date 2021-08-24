package main

import (
	"fmt"
	"sort"
	"strings"
)

// === Exercises === //

// ---------------------------------------------------------
// EXERCISE: Declare nil slices
//
//  1. Declare the following slices as nil slices:
//    1. The names of your friends (names slice)
//    2. The distances to locations (distances slice)
//    3. A data buffer (data slice)
//    4. Currency exchange ratios (ratios slice)
//    5. Up/Down status of web servers (alives slice)
//
//  2. Print their type, length and whether they're equal to nil value or not.
//
// EXPECTED OUTPUT
//  names    : []string 0 true
//  distances: []int 0 true
//  data     : []uint8 0 true
//  ratios   : []float64 0 true
//  alives   : []bool 0 true
// ---------------------------------------------------------

// func main() {
// 	friends := []string{}
// 	distances := []int{}
// 	data := []uint8{}
// 	ratios := []float64{}
// 	alives := []bool{}

// 	fmt.Printf("friends   : %T %d %t\n", friends, len(friends), friends == nil)
// 	fmt.Printf("distances : %T %d %t\n", distances, len(distances), distances == nil)
// 	fmt.Printf("data      : %T %d %t\n", data, len(data), data == nil)
// 	fmt.Printf("ratios    : %T %d %t\n", ratios, len(ratios), ratios == nil)
// 	fmt.Printf("alives    : %T %d %t\n", alives, len(alives), alives == nil)
// }

// ---------------------------------------------------------
// EXERCISE: Fix the problems
//
//  1. Uncomment the code
//
//  2. Fix the problems
//
//  3. BONUS: Simplify the code
//
//
// EXPECTED OUTPUT
//   "Einstein and Shepard and Tesla"
//   ["Fire" "Kafka's Revenge" "Stay Golden"]
//   [1 2 3 5 6 7 8 9]
// ---------------------------------------------------------

// func main() {
// 	names := []string{
// 		"Einstein",
// 		"Shepard",
// 		"Tesla",
// 	}

// 	books := []string{
// 		"Stay Golden",
// 		"Fire",
// 		"Kafka's Revenge",
// 	}

// 	sort.Strings(books)

// 	// this time, do not change the nums array to a slice
// 	nums := [...]int{5, 1, 7, 3, 8, 2, 6, 9}

// 	// use the slicing expression to change the nums array to a slice below
// 	sort.Ints(nums[:])

// 	// Here: Use the strings.Join function to join the names (see the expected output)
// 	fmt.Printf("%q\n", strings.Join(names, " and "))

// 	fmt.Printf("%q\n", books)
// 	fmt.Printf("%d\n", nums)
// }

// ---------------------------------------------------------
// EXERCISE: Compare the slices
//  1. Split the namesA string and get a slice
//  2. Sort all the slices
//  3. Compare whether the slices are equal or not
//
// EXPECTED OUTPUT
//   They are equal.
//
// HINTS
//   1. strings.Split function splits a string and
//      returns a string slice
//   2. Comparing slices: First check whether their length
//      are the same or not; only then compare them.
//
// ---------------------------------------------------------

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	aSlice := strings.Split(namesA, ", ")
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	var eq string

	if len(aSlice) != len(namesB) {
		eq = "not "
	} else {
		sort.Strings(aSlice)
		sort.Strings(namesB)
		for i, name := range aSlice {
			if name != namesB[i] {
				eq = "not "
				break
			}
		}
	}
	fmt.Printf("They are %sequal.\n", eq)
}
