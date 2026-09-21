func calPoints(operations []string) int {
	stack := make([]int, len(operations))
	var sum int
	for _, s := range operations {
		switch s {
		case "+":
			newElement := stack[len(stack)-2] + stack[len(stack)-1]
			stack = append(stack, newElement)
			sum += newElement
		case "D":
			newElement := stack[len(stack)-1] * 2
			sum += newElement
			stack = append(stack, newElement)
		case "C":
			sum -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			newElement, _ := strconv.Atoi(s)
			sum += newElement
			stack = append(stack, newElement)
		}
	}
	return sum
}