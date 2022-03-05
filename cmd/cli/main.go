package main

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
