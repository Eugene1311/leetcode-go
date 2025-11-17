package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slowPointer := head
	fastPointer := head

	for fastPointer != nil {
		slowPointer = slowPointer.Next

		if fastPointer.Next == nil {
			return false
		}
		fastPointer = fastPointer.Next.Next

		if slowPointer == fastPointer {
			return true
		}
	}

	return false
}
