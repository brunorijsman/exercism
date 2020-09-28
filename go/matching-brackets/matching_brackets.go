// Go exercise matching-brackets
package brackets

var closeToOpen map[rune]rune = map[rune]rune{']': '[', '}': '{', ')': '('}

// Determine if string str contains matching brackets.
func Bracket(str string) bool {
	stack := []rune{}
	for _, r := range str {
		switch r {
		case '[', '{', '(':
			stack = append(stack, r)
		case ']', '}', ')':
			if len(stack) == 0 {
				return false
			}
			stackR := stack[len(stack)-1]
			if stackR != closeToOpen[r] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
		}
	}
	return len(stack) == 0
}
