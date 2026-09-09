func twoSum(nums []int, target int) []int {
    numIndxs := make(map[int]int, len(nums))

	for i, v := range nums {
		numIndxs[v] = i
	}
	output := []int{}

	for i, v := range nums {
		idx, ok := numIndxs[target-v]
		if ok && i != idx {
			output = append(output, i, idx)
			break
		}
	}

	return output
}
