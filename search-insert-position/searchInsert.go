package search_insert_position

import "fmt"

func searchInsert(nums []int, target int) int {
	return searchInsertHelper(nums, target, 0)
}

func searchInsertHelper(nums []int, target int, offset int) int {
	if len(nums) == 1 {
		value := nums[0]
		if value == target {
			return offset
		} else if value > target {
			return offset
		} else if value < target {
			return offset + 1
		}
	}

	middleIdx := len(nums) / 2

	middle := nums[middleIdx]
	if middle == target {
		return middleIdx + offset
	} else if middle > target {
		return searchInsertHelper(nums[:middleIdx], target, 0+offset)
	} else if middle < target {
		return searchInsertHelper(nums[middleIdx:], target, middleIdx+offset)
	}

	return -1
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}
