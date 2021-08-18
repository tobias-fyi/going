package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func sqrt(x float64) (float64, error) {
	if x < 0.0 { // Negative input throws an error
		return 0, ErrNegativeSqrt(x)
	}

	t, z := 0., 1.
	for {
		z, t = z-(z*z-x)/(2*z), z
		if math.Abs(t-z) < 1e-8 {
			break
		}
	}
	return z, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number.")
	}
	inum, err1 := strconv.ParseFloat(os.Args[1], 64)
	if err1 != nil {
		fmt.Printf("%s - Please provide valid number.", err1)
	}

	mathPkg := math.Sqrt(inum)
	newton, err2 := sqrt(inum)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("Number: %v\nNewton: %g\nmathPkg: %g\nDelta: %g\n",
		inum, newton, mathPkg, math.Abs(newton-mathPkg),
	)
}
