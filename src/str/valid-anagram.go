package str

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := make(map[uint8]int8)
	for i := 0; i < len(s); i++ {
		counts[s[i]] += 1
		counts[t[i]] -= 1
	}
	for _, v := range counts {
		if v < 0 {
			return false
		}
	}
	return true
}
