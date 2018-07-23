package kyu_5

var m = map[string]string{
	"{": "}",
	"[": "]",
	"(": ")",
	"<": ">",
}

func StackValidParentheses(str string) bool {
	stack := make([]string, 0)
	for _, value := range str {
		value := string(value)
		if len(stack) != 0 && m[stack[len(stack)-1]] == value {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, value)
		}
	}

	return len(stack) == 0
}
