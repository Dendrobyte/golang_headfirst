package chapters

import "fmt"

// Pool Puzzle
type Population int

// Exercise pg 278
type Number int

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func Chapter9() {
	// Pool Puzzle
	var population Population
	population = Population(573)
	fmt.Println("Sleepy Creek County population:", population)
	fmt.Println("Congratulations, Keven and Ann! It's a girl!")
	population += 1
	fmt.Println("SleepyCreekCounty population:", population)

	// Exercise pg 278
	ten := Number(10)
	ten.Add(4)
	ten.Subtract(5)
	four := Number(4)
	four.Add(3)
	four.Subtract(2)
}
