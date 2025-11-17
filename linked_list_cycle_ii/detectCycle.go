package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slowPointer := head
	fastPointer := head

	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next

		if slowPointer == fastPointer {
			break
		}
	}

	if fastPointer == nil || fastPointer.Next == nil {
		return nil
	}

	slowPointer = head
	for slowPointer != fastPointer {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next
	}
	return slowPointer
}
