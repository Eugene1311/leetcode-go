package main

import (
	"fmt"
	"math"
)

func maxScore(nums []int, x int) int64 {
	if len(nums) == 0 {
		return 0
	}

	scores := make([]int, 0)
	currentScore := nums[0]
	currentParity := nums[0] % 2
	for i := 1; i < len(nums); i++ {
		if nums[i]%2 == currentParity {
			currentScore += nums[i]
		} else {
			scores = append(scores, currentScore)
			currentScore = nums[i]
			currentParity = nums[i] % 2
		}
	}
	scores = append(scores, currentScore)
	//fmt.Println(scores)
	if len(scores) == 1 {
		return int64(scores[0])
	}

	dp := make([]int, len(scores))
	dp[0] = scores[0]
	dp[1] = scores[0] + scores[1] - x
	for i := 2; i < len(scores); i++ {
		dp[i] = int(math.Max(float64(dp[i-2]+scores[i]), float64(dp[i-1]+scores[i]-x)))
	}
	//fmt.Println(dp)
	var result int64 = 0
	for _, d := range dp {
		result = int64(math.Max(float64(result), float64(d)))
	}

	return result
}

func main() {
	fmt.Println(maxScore([]int{2, 3, 6, 1, 9, 2}, 5))
	fmt.Println(maxScore([]int{2, 4, 6, 8}, 3))
	fmt.Println(maxScore([]int{38, 92, 23, 30, 25, 96, 6, 71, 78, 77, 33, 23, 71, 48, 87, 77, 53, 28, 6, 20, 90, 83, 42, 21, 64, 95, 84, 29, 22, 21, 33, 36, 53, 51, 85, 25, 80, 56, 71, 69, 5, 21, 4, 84, 28, 16, 65, 7}, 52)) // 1545
	fmt.Println(maxScore([]int{9, 58, 17, 54, 91, 90, 32, 6, 13, 67, 24, 80, 8, 56, 29, 66, 85, 38, 45, 13, 20, 73, 16, 98, 28, 56, 23, 2, 47, 85, 11, 97, 72, 2, 28, 52, 33}, 90))                                             // 886
}
