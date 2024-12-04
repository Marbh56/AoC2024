package main

import (
	"fmt"
	"os"
	"strings"
)

func rotateGrid(grid []string) []string {
	// Grid dimensions will be implicitly derived from the grid itself.
	numRows := len(grid)
	numCols := len(grid[0])

	// Create a new grid for the rotated version
	rotatedGrid := make([]string, numCols)
	for i := 0; i < numCols; i++ {
		var sb strings.Builder
		for j := numRows - 1; j >= 0; j-- {
			sb.WriteByte(grid[j][i])
		}
		rotatedGrid[i] = sb.String()
	}
	return rotatedGrid
}

func matchPattern(grid []string, pattern []string) int {
	// No need to declare numRows and numCols globally, we can derive them here
	numRows := len(grid)
	numCols := len(grid[0])

	totalMatches := 0

	for i := 0; i < numRows-2; i++ { // Ensure no out-of-bounds errors
		for j := 0; j < numCols-2; j++ { // Ensure no out-of-bounds errors
			// Check for the specific M.S, .A., M.S pattern
			if grid[i][j] == pattern[0][0] && grid[i][j+2] == pattern[0][2] &&
				grid[i+1][j+1] == pattern[1][1] &&
				grid[i+2][j] == pattern[2][0] && grid[i+2][j+2] == pattern[2][2] {
				totalMatches++
			}
		}
	}
	return totalMatches
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("problem opening file:", err)
		return
	}
	// Read the input as rows of a grid
	str := string(file)
	rows := strings.Split(strings.TrimSpace(str), "\n")

	// The pattern to look for: M.S, .A., M.S
	pattern := []string{
		"M.S",
		".A.",
		"M.S",
	}

	totalMatches := 0

	// Rotate the grid 4 times
	for i := 0; i < 4; i++ {
		// Count matches for the current rotation
		totalMatches += matchPattern(rows, pattern)

		// Rotate the grid for the next iteration
		rows = rotateGrid(rows)
	}

	fmt.Println("Total Matches:", totalMatches)
}
