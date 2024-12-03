package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var safeReports = 0
	content, err := os.ReadFile("input.txt")
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

		if (isAscending(report) || isDescending(report)) &&
			differencesWithinLimit(report, 1, 3) {
			safeReports++
			fmt.Printf("Report is safe: %v\n", report)
		} else {
			fmt.Printf("Report is not safe: %v\n", report)
			if problemDampener(report) {
				safeReports++
				fmt.Printf("Problem Dampener marked the report as safe: %v\n", report)
			} else {
				fmt.Printf("Problem Dampener failed to mark the report as safe: %v\n", report)
			}
		}
		fmt.Println("--------------------------------------------------")
	}
	fmt.Printf("There are %v safe reports\n", safeReports)
}

func combiGenerator(report []int) [][]int {
	subReports := make([][]int, 0) // Initialize an empty slice of slices for subReports

	for i := 0; i < len(report); i++ {
		// Create a sub-report by concatenating slices before and after index i
		combo := append([]int(nil), report[:i]...) // Copy elements before i
		combo = append(combo, report[i+1:]...)     // Append elements after i
		subReports = append(subReports, combo)     // Add the resulting sub-report to the list
	}

	return subReports
}

func problemDampener(report []int) bool {
	fmt.Printf("Problem Dampener is checking: %v\n", report)
	subReports := combiGenerator(report)

	for _, subReport := range subReports {
		if (isAscending(subReport) || isDescending(subReport)) &&
			differencesWithinLimit(subReport, 1, 3) {
			return true // Return true immediately when a valid sub-report is found
		}
	}
	return false // Return false if no valid sub-report is found after checking all
}

func isAscending(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if report[i] >= report[i+1] {
			return false
		}
	}
	return true
}

func isDescending(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if report[i] <= report[i+1] {
			return false
		}
	}
	return true
}

func differencesWithinLimit(report []int, minDiff, maxDiff int) bool {
	for i := 0; i < len(report)-1; i++ {
		diff := abs(report[i] - report[i+1])
		if diff < minDiff || diff > maxDiff {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
