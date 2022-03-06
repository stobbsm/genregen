package list

import (
	_ "embed"
	"errors"
	"math/rand"
	"strings"
	"time"
)

// List of phrases that can be joined
// © Matthew Stobbs, 2022

//go:embed builtin.list
var dw []byte

var words = parse(dw)

// GetRandom returns a random word from the word list
func GetRandom() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(words))
	return words[n]
}

// GetRandomN returns the number of random strings defined in n int as a slice of size n
// if n > len(words), it returns nil and an error instead
func GetRandomN(n int) ([]string, error) {
	if n > len(words) {
		return nil, errors.New("asked to get more words then the list contains")
	}
	var t string
	w := make([]string, n)
	for i := 0; i < n; {
		t = GetRandom()
		if i == 0 {
			w[i] = t
			i++
			continue
		} else if sliceContains(w, t) {
			continue
		}
		w[i] = t
		i++
	}
	return w, nil
}

// sliceContains returns true if the slice contains the string, false if not
func sliceContains(ss []string, w string) bool {
	for _, s := range ss {
		if strings.Compare(s, w) == 0 {
			return true
		}
	}
	return false
}
