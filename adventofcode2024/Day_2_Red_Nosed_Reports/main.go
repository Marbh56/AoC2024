package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Safe if all increasing or all decreasing
	//Any two adjacent levels differ by at least one and at most three.
    // Need to adjust to only counting failures for isAscending or isDescending not bot
	var safeReports = 0
	content, err := os.ReadFile("test_input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		var report []int
		for _, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				log.Fatalf("error converting to int: %v", err)
			}
			report = append(report, number)
		}
		if isReportSafe(report) {
			fmt.Printf("Report %v is safe\n", report)
			safeReports++
		} else {
			fmt.Printf("Report %v is not safe\n", report)
		}
	}
	fmt.Printf("There are %v safe reports\n", safeReports)
}

func isReportSafe(numbers []int) bool {
	ascFails := isAscending(numbers)
	descFails := isDescending(numbers)
	diffFails := differencesWithinLimit(numbers, 1, 3)
	totalFailures := ascFails + descFails + diffFails
	return totalFailures <= 1
}

func isAscending(numbers []int) int {
	failures := 0
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] >= numbers[i+1] {
			failures++
		}
	}
	return failures

}

func isDescending(numbers []int) int {
	failures := 0
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] <= numbers[i+1] {
			failures++
		}
	}
	return failures
}

func differencesWithinLimit(numbers []int, minDiff, maxDiff int) int {
	failures := 0
	for i := 0; i < len(numbers)-1; i++ {
		diff := abs(numbers[i] - numbers[i+1])
		if diff < minDiff || diff > maxDiff {
			failures++
		}
	}
	return failures

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
