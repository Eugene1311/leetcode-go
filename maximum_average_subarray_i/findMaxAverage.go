package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := float64(sum)

	for i := k; i < len(nums); i++ {
		sum -= nums[i-k]
		sum += nums[i]
		maxSum = math.Max(maxSum, float64(sum))
	}

	return maxSum / float64(k)
}
