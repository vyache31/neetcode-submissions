func firstMissingPositive(nums []int) int {
	inNums := make(map[int]bool)
	var max int
	for _, v := range nums {
		inNums[v] = true
		if v > max {
			max = v
		}
	}
	for num := range max {
		if num != 0 && inNums[num] == false {
			return num
		}
	}
	return max + 1
}