func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := map[rune]int{}
	for _, c := range s {
		counts[c]++
	}

	for _, c := range t {
		counts[c]--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true

}
