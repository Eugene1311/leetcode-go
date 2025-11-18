package main

import "testing"

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{name: "test 1", nums: []int{1, 3, 4, 2, 2}, expected: 2},
		{name: "test 2", nums: []int{3, 1, 3, 4, 2}, expected: 3},
		{name: "test 3", nums: []int{3, 3, 3, 3, 3}, expected: 3},
		{name: "test 4", nums: []int{18, 13, 14, 17, 9, 19, 7, 17, 4, 6, 17, 5, 11, 10, 2, 15, 8, 12, 16, 17}, expected: 17},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.name)
			actual := findDuplicate(test.nums)
			if test.expected != actual {
				t.Errorf("wrong result: %d for %+v, expected: %d", actual, test.nums, test.expected)
			}
		})
	}
}
