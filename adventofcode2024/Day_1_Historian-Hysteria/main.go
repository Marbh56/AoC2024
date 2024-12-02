package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var leftNums, rightNums []int
	var simScore = 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	text := string(content)
	numbers := strings.Fields(text)
	for i, number := range numbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			log.Fatalf("failed to convert to int: %s", err)
		}
		if i%2 == 0 {
			leftNums = append(leftNums, num)
		} else {
			rightNums = append(rightNums, num)
		}
	}

	sort.Ints(leftNums)
	sort.Ints(rightNums)

	var distances []int
	for i := range leftNums {
		if i < len(rightNums) {
			diff := leftNums[i] - rightNums[i]
			if diff < 0 {
				diff = -diff
			}
			distances = append(distances, diff)
		}
	}
	totalDistance := 0
	for _, d := range distances {
		totalDistance += d
	}

	for _, leftNum := range leftNums {
		count := 0
		for _, rightNum := range rightNums {
			if leftNum == rightNum {
				count++
			}
		}
		simScore += leftNum * count
	}

	fmt.Printf("Total distance: %d\n", totalDistance)
	fmt.Printf("Sim score: %d\n", simScore)

}
