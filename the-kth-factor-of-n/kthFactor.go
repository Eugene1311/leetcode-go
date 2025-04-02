package the_kth_factor_of_n

func kthFactor(n int, k int) int {
	factors := make([]int, 0)

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			if len(factors) == k {
				return factors[k-1]
			}
		}
	}

	return -1
}
