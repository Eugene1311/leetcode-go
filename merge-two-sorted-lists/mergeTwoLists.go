package merge_two_sorted_lists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var head, currentNode *ListNode
	nextNode1 := list1
	nextNode2 := list2
	if list1.Val <= list2.Val {
		currentNode = list1
		nextNode1 = list1.Next
	} else {
		currentNode = list2
		nextNode2 = list2.Next
	}
	head = currentNode

	for nextNode1 != nil && nextNode2 != nil {
		if nextNode1.Val <= nextNode2.Val {
			currentNode.Next = nextNode1
			nextNode1 = nextNode1.Next
		} else {
			currentNode.Next = nextNode2
			nextNode2 = nextNode2.Next
		}

		currentNode = currentNode.Next
	}

	if nextNode1 == nil {
		currentNode.Next = nextNode2
	} else if nextNode2 == nil {
		currentNode.Next = nextNode1
	}

	return head
}

func main() {
	list1 := getListFromArray([]int{1, 3, 5})
	list2 := getListFromArray([]int{2, 4, 6})
	mergedList := mergeTwoLists(list1, list2)
	fmt.Println(mergedList.Val)
}

func getListFromArray(arr []int) *ListNode {
	var head *ListNode
	var currentNode *ListNode
	for i, item := range arr {
		nextItem := &ListNode{
			Val: item,
		}
		if i == 0 {
			head = nextItem
			currentNode = head
		} else {
			currentNode.Next = nextItem
			currentNode = currentNode.Next
		}
	}

	return head
}
