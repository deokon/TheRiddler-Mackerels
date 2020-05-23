package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type wordList struct {
	loaded bool
	words  map[string]int
}

// readLines reads a whole file into memory and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func loadDictionary(f string) wordList {
	var dict = wordList{loaded: false, words: map[string]int{}}

	// populate dictionary
	lines, err := readLines(f)
	if err != nil {
		log.Fatal("Could not read dictionary file\n")
	}

	// build dictionary with word scores
	for _, w := range lines {
		dict.words[w] = 0
	}
	dict.loaded = true

	fmt.Printf("Loaded dictionary of %d words\n", len(dict.words))

	return dict
}

// InDictionary returns true if the string is in the dictionary file and false otherwise
func InDictionary(s string, dict wordList) bool {
	_, ok := dict.words[strings.ToLower(s)]
	return ok
}
