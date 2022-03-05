package main

// CLI application for generating book genre's.
// © Matthew Stobbs, 2022

import (
	"fmt"
	"strings"

	words "github.com/stobbsm/genregen"
)

func main() {
	wordList, err := words.GetRandomN(2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated: %s\n", strings.Join(wordList, " "))
}
