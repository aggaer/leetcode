package str

func CheckInclusion(s1 string, s2 string) bool {
	l, r, match, capP := 0, 0, 0, 0
	need := make([]int, 26)
	window := make([]int, 26)
	for _, v := range s1 {
		if need[v-'a'] == 0 {
			capP++
		}
		need[v-'a']++
	}
	for r < len(s2) {
		cur := s2[r] - 'a'
		if need[cur] != 0 {
			window[cur]++
			if window[cur] == need[cur] {
				match++
			}
		}
		r++
		//若窗口匹配则l向右缩进，直到满足条件
		for match == capP {
			if r-l == len(s1) {
				return true
			}
			//l向左缩进到r的位置
			cur := s2[l] - 'a'
			if window[cur] != 0 {
				window[cur]--
				if window[cur] < need[cur] {
					match--
				}
			}
			l++
		}
	}
	return false
}
