package maximum_product_of_two_digits

import (
	"sort"
)

func maxProduct(n int) int {
	digits := make([]int, 0)
	delimiter := 10

	for n > 0 {
		digit := n % delimiter
		digits = append(digits, digit)
		n /= 10
	}

	sort.Ints(digits)

	return digits[len(digits)-1] * digits[len(digits)-2]
}
