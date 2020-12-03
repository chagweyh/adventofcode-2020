package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var expenses []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		expense, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}

func getAnswer(expenses []int) int {
	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses)-1; j++ {
			if expenses[i]+expenses[j] == 2020 {
				return expenses[i] * expenses[j]
			}
		}
	}

	return 0
}

func main() {
	expenses, err := parseFile("expenses.txt")
	if err != nil {
		panic(err)
	}

	answer := getAnswer(expenses)
	fmt.Println(answer)
}
