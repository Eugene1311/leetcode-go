package valid_parentheses

var openingToClosingParentheses = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}
var closingToOpeningParentheses = map[rune]rune{
	'}': '{',
	')': '(',
	']': '[',
}

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, ch := range s {
		_, ok := openingToClosingParentheses[ch]
		if ok {
			stack = append(stack, ch)
		} else if len(stack) > 0 && stack[len(stack)-1] == closingToOpeningParentheses[ch] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}
