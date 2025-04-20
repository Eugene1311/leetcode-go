package minimum_operations_to_make_array_sum_divisible_by_k

func minOperations(nums []int, k int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum % k
}
