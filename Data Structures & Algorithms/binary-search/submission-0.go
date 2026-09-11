func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for {
		if l > r {
			break
		}
		m := l + (r-l)/2
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