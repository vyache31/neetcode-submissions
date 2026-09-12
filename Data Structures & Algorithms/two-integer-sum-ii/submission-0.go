func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for {
		sum := numbers[l] + numbers[r]
		if sum < target {
			l += 1
		} else if sum > target {
			r -= 1
		} else {
			break
		}
	}
	return []int{ l+1, r+1 }
}