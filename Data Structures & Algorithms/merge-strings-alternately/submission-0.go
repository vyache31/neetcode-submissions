func mergeAlternately(word1 string, word2 string) string {
	out := make([]byte, len(word1)+len(word2))
	var p1, p2, outIdx int
	for p1 < len(word1) && p2 < len(word2) {
		out[outIdx], out[outIdx+1] = word1[p1], word2[p2]
		outIdx += 2
		p1++
		p2++
	}
	for p1 < len(word1) {
		out[outIdx] = word1[p1]
		outIdx++
		p1++
	}
	for p2 < len(word2) {
		out[outIdx] = word2[p2]
		outIdx++
		p2++
	}
	return string(out)
}
