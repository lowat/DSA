package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct{
		input []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
        	{[]int{1, 2, 3, 4}, false},
        	{[]int{1,1,1,3,3,4,3,2,4,2}, true},
        	{[]int{}, false},
        	{[]int{5}, false},
	}

	for _, test := range tests {
		result := ContainsDuplicate(test.input)
		if result != test.expected {
			t.Errorf("ContainsDuplicate(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
