package main

import (
	"bufio"
	"fmt"
	"strings"
)

func Prompt(scanner *bufio.Scanner) []string {
	fmt.Print("Pokedex > ")
	scanner.Scan()
	return cleanInput(scanner.Text())
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
