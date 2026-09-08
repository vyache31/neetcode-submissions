func isAnagram(s string, t string) bool {
	sCnts := make(map[rune]int, len(s))
	tCnts := make(map[rune]int, len(s))
	for _, symbol := range s {
		sCnts[symbol]++
	}
	for _, symbol := range t {
		tCnts[symbol]++
	}
	for key, value := range sCnts {
		if tCnts[key] != value {
			return false
		}
	}
	for key, value := range tCnts {
		if sCnts[key] != value {
			return false
		}
	}
	return true
}
