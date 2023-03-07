package chapters

import "fmt"

func Chapter6() {

	// Pool puzzle
	numbers := make([]float64, 3)
	numbers[0] = 19.7
	numbers[2] = 25.2
	for i, number := range numbers {
		fmt.Println(i, number)
	}

	var letters = []string{"a", "b", "c"}
	for i, letter := range letters {
		fmt.Println(i, letter)
	}

	// Code Magnets
	fmt.Println(sum(9, 7))
	fmt.Println(sum(1, 2, 4))

}

// Code Magnets
func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
