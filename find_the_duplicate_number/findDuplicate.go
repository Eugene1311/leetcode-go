package main

func findDuplicate(nums []int) int {
	slowPointer := 0
	fastPointer := 0

	for {
		slowPointer = nums[slowPointer]
		fastPointer = nums[nums[fastPointer]]

		if slowPointer == fastPointer {
			break
		}
	}

	slowPointer = 0

	for slowPointer != fastPointer {
		slowPointer = nums[slowPointer]
		fastPointer = nums[fastPointer]
	}

	return slowPointer
}
