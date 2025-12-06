package testdata

import "efedoseev/leetcode/reverse-linked-list-ii/lists"

func BuildList() *lists.ListNode {
	Node7 := &lists.ListNode{
		Val:  7,
		Next: nil,
	}
	Node6 := &lists.ListNode{
		Val:  6,
		Next: Node7,
	}
	Node5 := &lists.ListNode{
		Val:  5,
		Next: Node6,
	}
	Node4 := &lists.ListNode{
		Val:  4,
		Next: Node5,
	}
	Node3 := &lists.ListNode{
		Val:  3,
		Next: Node4,
	}
	Node2 := &lists.ListNode{
		Val:  2,
		Next: Node3,
	}
	Node1 := &lists.ListNode{
		Val:  1,
		Next: Node2,
	}

	return Node1
}
