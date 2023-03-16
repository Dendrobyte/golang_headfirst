package chapters

import (
	"fmt"
)

// Code Magnets
func repeat(s string, channel chan string) {
	for i := 0; i < 25; i++ {
		channel <- s
	}
}

// Code Magnets
func Chapter13() {

	testChan := make(chan string)
	go repeat("x", testChan)
	go repeat("y", testChan)
	for i := 0; i < 50; i++ {
		// Setting this to > 50 results in a deadlock, so go knows the channel will never update. Cool!
		fmt.Println(<-testChan)
	}

}
