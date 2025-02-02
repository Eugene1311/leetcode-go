package find_subarrays_with_equal_sum

type set map[int]struct{}

func findSubArrays(nums []int) bool {
	sums := make(set)
	for i := 0; i < len(nums)-1; i++ {
		sum := nums[i] + nums[i+1]
		if _, ok := sums[sum]; ok {
			return true
		}
		sums[sum] = struct{}{}
	}

	return false
}
