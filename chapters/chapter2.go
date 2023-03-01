package chapters

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func Chapter2() {
	fmt.Println("Chapter 2 Exercises")

	/* Code Magnets, page 43 */
	fileInfo, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal("Fileinfo has error")
	}
	fmt.Println(fileInfo.Size())
}

/* Number Guesser */
func Ch2NumGuesser() {
	rand.Seed(time.Now().Unix()) // Seed the randomness based in unix time
	target := rand.Intn(100) + 1
	fmt.Println("I've generated a number between 1 and 100... guess away!")
	reader := bufio.NewReader(os.Stdin)
	var x int
	success := false
	for x = 0; x <= 10; x++ {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("The input errored out!")
		}
		input = strings.TrimSpace(input)
		inputNum, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal("The input didn't seem acceptable... (whole numbers only)")
		}
		if inputNum < target {
			fmt.Println("Your guess was too LOW!")
		} else if inputNum > target {
			fmt.Println("Your guess was too HIGH!")
		} else {
			fmt.Println("Winner winner chicken dinner!")
			success = true
			break
		}
	}
	// Interestingly, inputNum is 0 here but is correct when we're in the above else block. Hmm.
	if success {
		fmt.Println("Congrats for guessing! It took you", x+1, "guesses.")
	} else {
		fmt.Println("Looks like you ran out of guesses, oh well! The answer was", target)
	}
}
