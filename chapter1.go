package main

import "fmt"

func main() {
	fmt.Println("Chapter 1 Exercises")
	
	/* Code Magnets, page 18 */
	var originalCount int = 10
	var eatenCount int = 4

	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "left.")

}
