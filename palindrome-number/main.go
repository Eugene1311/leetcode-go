package palindrome_number

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var reversed int
	y := x
	for y > 0 {
		digit := y % 10
		reversed = reversed*10 + digit
		y = y / 10
	}

	return reversed == x
}

func main() {
	x := 121
	fmt.Println(isPalindrome(x))
}
