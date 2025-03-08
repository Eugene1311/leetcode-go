package range_sum_query_immutable

import "fmt"

type NumArray struct {
	nums       []int
	prefixSums []int
}

func Constructor(nums []int) NumArray {
	prefixSums := make([]int, len(nums)+1)
	for i := -1; i < len(nums); i++ {
		if i == -1 {
			prefixSums[i+1] = 0
		} else {
			prefixSums[i+1] = prefixSums[i] + nums[i]
		}
	}
	return NumArray{nums, prefixSums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSums[right+1] - this.prefixSums[left]
}

func main() {
	obj := Constructor([]int{1, 2, 3, 4, 5})
	param1 := obj.SumRange(2, 4)
	fmt.Println(param1)
}
