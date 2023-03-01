package chapters

import (
	"errors"
	"fmt"
)

// Pool puzzle
func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return dividend / divisor, nil
}

func Chapter3() {
	// Pool puzzle
	quotient, err := divide(5.6, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}

	// Code Magnets pg 106
	var myInt int
	var myIntPointer *int
	myInt = 42
	myIntPointer = &myInt
	fmt.Println(*myIntPointer)
}

// Next chapter is packages, but I'll end on a good note today...
