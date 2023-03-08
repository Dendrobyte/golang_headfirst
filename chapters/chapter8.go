package chapters

import "fmt"

// Code Magnets
type student struct {
	name  string
	grade float64
}

// Pool Puzzle
type Coordinates struct {
	Longitude float64
	Latitude  float64
}

func printInfo(s *student) {
	fmt.Println("Name:", s.name)
	fmt.Printf("Grade: %0.1f\n", s.grade)
}

func Chapter8() {
	// Code Magnets
	var s student
	s.name = "Alonzo Cole"
	s.grade = 92.3
	printInfo(&s)

	// Pool Puzzle
	location := Coordinates{Latitude: 37.42, Longitude: -122.08}
	fmt.Println("Latitude:", location.Latitude)
	fmt.Println("Longitude:", location.Longitude)

}
