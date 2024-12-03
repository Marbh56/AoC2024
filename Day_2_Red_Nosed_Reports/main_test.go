package main

import (
	"reflect"
	"testing"
)

func TestCombiGenerator(t *testing.T) {
	testCases := []struct {
		report   []int
		expected [][]int
	}{
		{
			report:   []int{1, 2, 3},
			expected: [][]int{{2, 3}, {1, 3}, {1, 2}},
		},
		{
			report:   []int{5, 6, 7, 8},
			expected: [][]int{{6, 7, 8}, {5, 7, 8}, {5, 6, 8}, {5, 6, 7}},
		},
		{
			report:   []int{1},
			expected: [][]int{},
		},
		{
			report:   []int{},
			expected: [][]int{},
		},
		{
			report:   []int{1, 2},
			expected: [][]int{{2}, {1}},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := combiGenerator(tc.report)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("combiGenerator(%v) = %v; expected %v", tc.report, result, tc.expected)
			}
		})
	}
}
