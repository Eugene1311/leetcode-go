package main

import (
	"efedoseev/leetcode/reverse-linked-list-ii/lists"
	"efedoseev/leetcode/reverse-linked-list-ii/testdata"
	"slices"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name     string
		head     *lists.ListNode
		left     int
		right    int
		expected []int
	}{
		{
			name:     "test 1",
			head:     testdata.BuildList(),
			left:     3,
			right:    5,
			expected: []int{1, 2, 5, 4, 3, 6, 7},
		},
		{
			name:     "test 2",
			head:     testdata.BuildList(),
			left:     3,
			right:    7,
			expected: []int{1, 2, 7, 6, 5, 4, 3},
		},
		{
			name: "test 3",
			head: &lists.ListNode{
				Val: 3,
				Next: &lists.ListNode{
					Val:  5,
					Next: nil,
				},
			},
			left:     1,
			right:    1,
			expected: []int{3, 5},
		},
		{
			name: "test 4",
			head: &lists.ListNode{
				Val:  3,
				Next: nil,
			},
			left:     1,
			right:    1,
			expected: []int{3},
		},
		{
			name: "test 4",
			head: &lists.ListNode{
				Val: 3,
				Next: &lists.ListNode{
					Val:  5,
					Next: nil,
				},
			},
			left:     1,
			right:    2,
			expected: []int{5, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := listToSlice(reverseBetween(test.head, test.left, test.right))
			if !slices.Equal(test.expected, actual) {
				t.Errorf("wrong result, expected: %v, actual: %v", test.expected, actual)
			}
		})
	}
}

func listToSlice(head *lists.ListNode) []int {
	current := head
	result := make([]int, 0)

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}
