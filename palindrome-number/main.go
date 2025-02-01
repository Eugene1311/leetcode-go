package palindrome_number

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var reversed int
	temp := x
	for temp > 0 {
		digit := temp % 10
		reversed = reversed*10 + digit
		temp = temp / 10
	}

	return reversed == x
}

func main() {
	x := 121
	fmt.Println(isPalindrome(x))
}
