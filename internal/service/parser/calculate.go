package parser

func calculate(s string) int {
	stack := []int{}
	curNumber := 0
	operator := '+'

	for i, ch := range s {
		if isNumber(ch) {
			curNumber = curNumber*10 + int(ch-'0')
		}
		if !isNumber(ch) && ch != ' ' || i == len(s)-1 {
			if operator == '+' {
				stack = append(stack, curNumber)
			}
			if operator == '-' {
				stack = append(stack, -curNumber)
			}
			if operator == '*' {
				left := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, left*curNumber)
			}
			if operator == '/' {
				left := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, left/curNumber)
			}
			operator = ch
			curNumber = 0
		}
	}
	result := 0
	for _, num := range stack {
		result += num
	}
	return result
}

func isNumber(ch rune) bool { return '0' <= ch && ch <= '9' }
