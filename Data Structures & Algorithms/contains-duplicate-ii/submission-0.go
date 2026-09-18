func containsNearbyDuplicate(nums []int, k int) bool {
	idxs := make(map[int]int)
	for i, num := range nums {
		j, ok := idxs[num]
		if !ok {
			idxs[num] = i
			continue
		} else if abs(i-j) <= k {
			return true
		}
		idxs[nums[i]] = i
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}