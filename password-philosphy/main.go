package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	min      int
	max      int
	letter   byte
	password string
}

func parseFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	nbValid := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry, err := parseLine(scanner.Text())
		if err != nil {
			return 0, err
		}

		if isValid(*entry) {
			nbValid++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return nbValid, nil
}

func parseLine(s string) (*Entry, error) {
	entry := &Entry{}
	_, err := fmt.Sscanf(s, "%d-%d %c: %s", &entry.min, &entry.max, &entry.letter, &entry.password)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func isValid(entry Entry) bool {
	count := strings.Count(entry.password, string(entry.letter))
	return entry.min <= count && count <= entry.max
}

func main() {
	nbValid, err := parseFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(nbValid)
}
