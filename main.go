package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		userInput := Prompt(scanner)
		if len(userInput) > 0 {
			fmt.Printf("Your command was: %v\n", userInput[0])
		}
	}
}
