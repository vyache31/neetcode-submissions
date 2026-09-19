func searchInsert(nums []int, target int) int {
    l := 0
	r := len(nums) - 1
	for {
        m := l + (r-l)/2
		if l > r {
			return m
		}
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}