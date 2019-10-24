package str

import "github.com/aggaer/leetcode/src/utils"

func lengthOfLongestSubstring(s string) int {
	start, end, result := 0, 0, 0
	m := make(map[rune]int)
	for _, c := range s {
		if pos, ok := m[c]; ok {
			start = utils.Max(start, pos)
		}
		result = utils.Max(result, end-start+1)
		m[c] = end + 1
		end++
	}
	return result
}

func main() {
	print(lengthOfLongestSubstring("abcabcbb"))
}
