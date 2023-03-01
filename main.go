package main

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

func main() {
	rand.Seed(time.Now().Unix()) // Seed the randomness based in unix time
	target := rand.Intn(100) + 1
	fmt.Println(target)
	fmt.Println("I've generated a number between 1 and 100... guess away!")
	reader := bufio.NewReader(os.Stdin)
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
	}
}
