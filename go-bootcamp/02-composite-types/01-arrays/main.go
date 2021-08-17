package main

import "fmt"

func main() {
	// Last two values are empty
	shop := [6]string{
		"tomatoes",
		"carrots",
		"celery",
		"hummus",
	}

	fmt.Printf("%#v\n", shop)

	// Auto sets length of array
	shop2 := [...]string{
		"tomatoes",
		"carrots",
		"celery",
		"hummus",
	}

	fmt.Printf("%#v\n", shop2)

	// Multidimensional arrays
	multi := [2][3]int{}

	fmt.Printf("%#v\n", multi)

	// Keyed elements
	rates := [...]float64{
		5:   1.5,
		2.5, // Inserted into [6]
		0:   0.5,
	}

	fmt.Printf("%#v\n", rates)

	// More keyed elements fun
	const (
		ETH = iota
		BTC
	)

	rates2 := [...]float64{
		ETH: 3158.69,
		BTC: 45975.36,
	}

	fmt.Printf("1 USD is %g ETH\n", rates2[ETH])
	fmt.Printf("1 USD is %g BTC\n", rates2[BTC])
}
