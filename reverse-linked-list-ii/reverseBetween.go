package main

import (
	"efedoseev/leetcode/reverse-linked-list-ii/lists"
)

// 1 -> 2 -> 3 -> 4 -> 5 -> 6 (3, 5)
// 1 -> 2 -> nil, 3 -> 4 -> 5 -> nil, 6
// 1 -> 2 -> nil, 5 -> 4 -> 3 -> nil, 6
// 1 -> 2 -> 5 -> 4 -> 3 -> 6
func reverseBetween(head *lists.ListNode, left int, right int) *lists.ListNode {
	current := head
	var preceding *lists.ListNode = nil

	i := 1
	for i < left {
		preceding = current
		current = current.Next
		i++
	}
	first := current

	var prev *lists.ListNode = nil
	var next *lists.ListNode = nil
	for i <= right {
		next = current.Next
		current.Next = prev
		prev = current
		current = next

		i++
	}
	if preceding != nil {
		preceding.Next = prev
	}
	first.Next = next

	if left == 1 {
		return prev
	}
	return head
}
