package chapters

import (
	"fmt"
	"log"
	"os"
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
