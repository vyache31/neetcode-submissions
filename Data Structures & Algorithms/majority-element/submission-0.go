func majorityElement(nums []int) int {
	cnt := make(map[int]int)
	var mjr int
	for _, v := range nums {
		cnt[v]++
		if cnt[v] > (len(nums) / 2) {
			mjr = v
			break
		}
	}
	return mjr
}
