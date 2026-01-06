package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
		{
			nums:     []int{-3, 4, 3, 90},
			target:   0,
			expected: []int{0, 2},
		},
		{
			nums:     []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3},
		},
		{
			nums:     []int{1, 5, 7, 2, 8, 11},
			target:   9,
			expected: []int{1, 3},
		},
		{
			nums:     []int{1000000, 500000, -500000},
			target:   500000,
			expected: []int{1, 2},
		},
		{
			nums:     []int{1, 2, 3, 4, 5, 9},
			target:   14,
			expected: []int{4, 5},
		},
		{
			nums:     []int{5, 9, 1, 2, 3},
			target:   14,
			expected: []int{0, 1},
		},
	}

	for _, test := range tests {
		actual := TwoSum(test.nums, test.target)

		sort.Ints(actual)
		sort.Ints(test.expected)

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf(
				"TwoSum(%v, %d) = %v; expected %v",
				test.nums, test.target, actual, test.expected,
			)
		}
	}
}

