package list

import (
	"os"
	"strings"
)

// Parse given file and return the string slice

// Default word list

// Open opens the given file, and returns the lines read as a []string.
// Each line is one entry.
// Return nil and an error if there is a problem reading the file or the contents
func Open(f string) ([]string, error) {
	b, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return parse(b), nil
}

func parse(b []byte) []string {
	return strings.Split(string(b), "\n")
}
