func isAnagram(s string, t string) bool {
	chrCnts := [26]int{} // size of eng alphabet is 26

	for i := 0; i < len(s); i++ {
		chrCnts[s[i]-97]++ // cause lowercase Eng letters in ASCII start at 97
	}
	for i := 0; i < len(t); i++ {
		chrCnts[t[i]-97]--
	}
	for _, value := range chrCnts {
		if value != 0 {
			return false
		}
	}
	return true
}
