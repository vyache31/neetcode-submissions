func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	var l int
	r := l + 1
	records := map[byte]int{s[l]: l}
	maxLength := r - l
	for r < len(s) {
		i, ok := records[s[r]]
		if ok && i >= l {
			maxLength = max(maxLength, r-l)
			l = i + 1
		} else {
			maxLength = max(maxLength, r-l+1)
		}
		records[s[r]] = r
		r++
	}
	return maxLength
}