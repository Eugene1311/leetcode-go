package main

import (
	"fmt"
	"math"
)

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	diff := goal - sum
	result := math.Abs(float64(diff / limit))

	if math.Mod(float64(diff), float64(limit)) != 0 {
		result++
	}

	return int(result)
}

func main() {
	fmt.Printf("result: %d\n", minElements([]int{1, -1, 1}, 3, -4))
	fmt.Printf("result: %d\n", minElements([]int{1, -10, 9, 1}, 100, 0))
	fmt.Printf("result: %d", minElements([]int{-1, 5, 1, 7, 6, -3}, 7, 977925843))
}
