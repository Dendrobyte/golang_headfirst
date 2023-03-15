package chapters

import (
	"fmt"
	"time"
)

// Code Magnets
func repeat(s string) {
	for i := 0; i < 25; i++ {
		fmt.Print(s)
	}
}

// Code Magnets
func Chapter13() {

	go repeat("x")
	go repeat("y")
	time.Sleep(time.Second)

}
