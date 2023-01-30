package parser

import "fmt"

func calculate(s string) (int, error) {
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
			} else if operator == '-' {
				stack = append(stack, -curNumber)
			} else if operator == '*' {
				left := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, left*curNumber)
			} else if operator == '/' {
				left := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if curNumber == 0 {
					return 0, fmt.Errorf("[Parcer.calculate]: division by zero")
				}
				stack = append(stack, left/curNumber)
			} else {
				return 0, fmt.Errorf("[Parcer.calculate]: incorrect operator")
			}
			operator = ch
			curNumber = 0
		}
	}
	result := 0
	for _, num := range stack {
		result += num
	}
	return result, nil
}

func isNumber(ch rune) bool { return '0' <= ch && ch <= '9' }
