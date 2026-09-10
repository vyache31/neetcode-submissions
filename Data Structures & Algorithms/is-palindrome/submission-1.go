func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	r := len(s) - 1
	for l := 0; l < len(s); l++ {
		if !isLetterOrNum(s[l]) {
			continue
		}
		for {
			if isLetterOrNum(s[r]) {
				break
			}
			r -= 1
		}
		if s[l] != s[r] {
			return false
		}
		if l == r {
			break
		}
		r -= 1
	}

	return true

}

func isLetterOrNum(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= '0' && b <= '9')
}