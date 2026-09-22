func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		switch token {
		case "+":
			result := stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		case "-":
			result := stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		case "/":
			result := stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		case "*":
			result := stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		default:
			num, err := strconv.Atoi(token)
			if err == nil {
				stack = append(stack, num)
			}
		}
	}
	return stack[len(stack)-1]
}