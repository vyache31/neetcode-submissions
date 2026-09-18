func containsNearbyDuplicate(nums []int, k int) bool {
	idxs := make(map[int]struct{}, k)
	for i, num := range nums {
		_, ok := idxs[num]
		if ok {
			return true
		}
		idxs[num] = struct{}{}
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