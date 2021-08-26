package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
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

// func main() {
// 	namesA := "Da Vinci, Wozniak, Carmack"
// 	aSlice := strings.Split(namesA, ", ")
// 	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

// 	var eq string

// 	if len(aSlice) != len(namesB) {
// 		eq = "not "
// 	} else {
// 		sort.Strings(aSlice)
// 		sort.Strings(namesB)
// 		for i, name := range aSlice {
// 			if name != namesB[i] {
// 				eq = "not "
// 				break
// 			}
// 		}
// 	}
// 	fmt.Printf("They are %sequal.\n", eq)
// }

// ---------------------------------------------------------
// EXERCISE: Fix the backing array problem
//
//  Ensure that changing the elements of the `mine` slice
//  does not change the elements of the `nums` slice.
//
//
// CURRENT OUTPUT (INCORRECT)
//
//  Mine         : [-50 -100 -150 25 30 50]
//  Original nums: [-50 -100 -150]
//
//
// EXPECTED OUTPUT
//
//  Mine         : [-50 -100 -150]
//  Original nums: [56 89 15]
//
// ---------------------------------------------------------

// func main() {
// 	// DON'T TOUCH THE FOLLOWING CODE
// 	nums := []int{56, 89, 15, 25, 30, 50}

// 	// ----------------------------------------
// 	// ONLY ADD YOUR CODE HERE
// 	// Ensure that nums slice never changes even though
// 	// the mine slice changes.
// // My solution
// 	mine := make([]int, len(nums))
// // Official solution
// mine := append([]int(nil), nums[:3]...)

// 	// DON'T TOUCH THE FOLLOWING CODE
// 	// This code changes the elements of the nums slice.
// 	mine[0], mine[1], mine[2] = -50, -100, -150

// 	fmt.Println("Mine         :", mine)
// 	fmt.Println("Original nums:", nums[:3])
// }

// ---------------------------------------------------------
// EXERCISE: Sort the backing array
//  1. Sort only the middle 3 items.
//  2. All the slices should see your changes.
//
// RESTRICTION
//  Do not sort manually. Sort by using the sort package.
//
// EXPECTED OUTPUT
//  Original: [... doom galaga frogger asteroids simcity ...]
//  Sorted  : [... doom galaga asteroids frogger simcity ...]
//
// HINT:
//   Middle items are         : [frogger asteroids simcity]
//   After sorting they become: [asteroids frogger simcity]
// ---------------------------------------------------------

// func main() {
// 	items := []string{
// 		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
// 		"asteroids", "simcity", "metroid", "defender", "rayman",
// 		"tempest", "ultima",
// 	}

// 	fmt.Println("Original:", items)
// 	sort.Strings(items[5:8])
// 	fmt.Println()
// 	fmt.Println("Sorted  :", items)
// }

// ---------------------------------------------------------
// EXERCISE: Observe the memory allocations
//
//  In this exercise, your goal is to observe the memory allocation
//  differences between arrays and slices.
//
//  You will create, assign arrays and slices then you will print
//  the memory usage of your program on each step.
//
//  Please follow the instructions inside the code.
//
//
// EXPECTED OUTPUT
//
//  Note that, your memory usage numbers may vary. However, the size of the
//  arrays and slices should be the same on your own system as well
//  (if you're on a 64-bit machine).
//
//
//  [initial memory usage]
//          > Memory Usage: 104 KB
//  [after declaring an array]
//          > Memory Usage: 78235 KB
//  [after copying the array]
//          > Memory Usage: 156365 KB
//  [inside passArray]
//          > Memory Usage: 234495 KB
//  [after slicings]
//          > Memory Usage: 234497 KB
//  [inside passSlice]
//          > Memory Usage: 234497 KB
//
//  Array's size : 80000000 bytes.
//  Array2's size: 80000000 bytes.
//  Slice1's size: 24 bytes.
//  Slice2's size: 24 bytes.
//  Slice3's size: 24 bytes.
//
//
// HINTS
//
//  I've declared a few functions to help you.
//
//    report function:
//    - Prints the memory usage.
//    - Just call it with a message that matches to the expected output.
//
//    passArray function:
//    - It automatically prints the memory usage.
//    - Accepts a [size]int array, so you can pass your array to it.
//
//    passSlice function:
//    - It automatically prints the memory usage.
//    - Accepts an int slice, so you can pass it one of your slices.
//
// ---------------------------------------------------------

const size = 1e7

func main() {
	// don't worry about this code.
	// it stops the garbage collector: prevents cleaning up the memory.
	// see the link if you're curious:
	// https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)
	debug.SetGCPercent(-1)

	// run the program to see the initial memory usage.
	report("initial memory usage")

	// 1. allocate an array with 10 million int elements
	//    the array's size will be equal to ~80MB
	//    hint: use the `size` constant above.
	var big [size]int
	// 2. print the memory usage (use the report func).
	report("after array declaration")

	// 3. copy the array to a new array.
	big2 := big
	// 4. print the memory usage
	report("after copying array")

	// 5. pass the array to the passArray function
	passArray(big)

	// 6. convert one of the arrays to a slice
	bigSlice := big2[:]

	// 7. slice only the first 1000 elements of the array
	bigOne := bigSlice[:1000]

	// 8. slice only the elements of the array between 1000 and 10000
	bigTwo := bigSlice[1000:10000]

	// 9. print the memory usage (report func)
	report("after slicings")

	// 10. pass the one of the slices to the passSlice function
	passSlice(bigTwo)

	// 11. print the sizes of the arrays and slices
	//     hint: use the unsafe.Sizeof function
	//     see more here: https://golang.org/pkg/unsafe/#Sizeof
	fmt.Println(unsafe.Sizeof(big))
	fmt.Println(unsafe.Sizeof(big2))
	fmt.Println(unsafe.Sizeof(bigSlice))
	fmt.Println(unsafe.Sizeof(bigOne))
	fmt.Println(unsafe.Sizeof(bigTwo))
}

// passes [size]int array — about 80MB!
//
// observe that passing an array to a function (or assigning it to a variable)
// affects the memory usage dramatically
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
//
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
// don't worry about this code
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
