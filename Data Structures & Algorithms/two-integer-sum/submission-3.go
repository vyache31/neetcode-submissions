func twoSum(nums []int, target int) []int {
    numIndxs := make(map[int]int, len(nums))
	output := make([]int, 2) // because output must have 2 elements

	for i, v := range nums {
		idx, ok := numIndxs[target-v]
		if ok {
			output[0] = idx
			output[1] = i
			break
		} else {
			numIndxs[v] = i
		}
	}

	return output
}