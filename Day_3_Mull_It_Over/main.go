package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Read the entire file into a byte slice
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("err reading file:", err)
		return
	}

	// Convert the byte slice to a string
	content := string(data)

	// Regular expression to find the patterns
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	total := 0
	processing := true

	// Scan through each character in the content
	for i := 0; i < len(content); {
		// Check for "don't", stop processing
		if strings.HasPrefix(content[i:], "don't") {
			processing = false
			i += len("don't")
		}

		// Check for "do", resume processing
		if strings.HasPrefix(content[i:], "do") {
			processing = true
			i += len("do")
		}

		if processing {
			// Get the substring from the current index
			substr := content[i:]

			// Find the first match in the current substring
			match := re.FindStringSubmatchIndex(substr)

			if match != nil && match[0] == 0 {
				// It's a valid mul() expression
				xStr, yStr := substr[match[2]:match[3]], substr[match[4]:match[5]]
				x, err1 := strconv.Atoi(xStr)
				y, err2 := strconv.Atoi(yStr)
				if err1 == nil && err2 == nil {
					result := x * y
					total += result
					fmt.Printf("Found match: %s, x * y = %d\n", substr[match[0]:match[1]], result)
				}
				// Move index past the current match
				i += match[1]
			} else {
				i++ // move to the next character
			}
		} else {
			i++ // move to the next character
		}
	}

	// Print total of all matches
	fmt.Println("Total:", total)
}
