package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("1.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}

	var elves []int
	idx := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Read the line
		line := scanner.Text()
		if line == "" {
			idx++
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("failed to parse line: %v", err)
		}
		if idx >= len(elves) {
			elves = append(elves, calories)
		} else {
			elves[idx] += calories
		}
	}

	// Part 1
	sort.Ints(elves)

	fmt.Printf("Highest: %d\n", elves[len(elves)-1])

	// Part 2

	// Find the sum of the top 3 elves
	var sum int
	for i := len(elves) - 1; i >= len(elves)-3; i-- {
		sum += elves[i]
	}

	fmt.Printf("Sum of top 3: %d\n", sum)
}
