func containsNearbyDuplicate(nums []int, k int) bool {
	idxs := make(map[int]int, k)
	for i, num := range nums {
		_, ok := idxs[num]
		if ok {
			return true
		}
		idxs[num] = i
		if len(idxs) > k {
			delete(idxs, nums[i-k])
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}