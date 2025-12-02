package reverse_linked_list

import "testing"

func TestReverseList(t *testing.T) {
	node5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	node4 := &ListNode{
		Val:  5,
		Next: node5,
	}
	node3 := &ListNode{
		Val:  5,
		Next: node4,
	}
	node2 := &ListNode{
		Val:  5,
		Next: node3,
	}
	node1 := &ListNode{
		Val:  5,
		Next: node2,
	}
	tests := []struct {
		name string
		head *ListNode
	}{
		{
			name: "test 1",
			head: node1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := reverseList(test.head)
			if actual != node5 && actual.Next != node4 {
				t.Errorf("wrong result, actual: %+v", actual)
			}
		})
	}
}
