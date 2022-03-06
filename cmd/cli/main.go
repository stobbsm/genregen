package main

// CLI application for generating book genre's.
// © Matthew Stobbs, 2022

import (
	"fmt"
	"strings"

	"github.com/stobbsm/genregen/list"
)

func main() {
	wordList, err := list.GetRandomN(2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated: %s\n", strings.Join(wordList, " "))
}
