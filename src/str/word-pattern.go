package str

import (
	"strings"
)

func WordPattern(pattern string, str string) bool {
	mp := make(map[string]byte)
	s := strings.Fields(str)
	if len(pattern) != len(s) {
		return false
	}
	for i, v := range s {
		if p, ok := mp[v]; ok {
			if pattern[i] != p {
				return false
			}
		} else {
			//没有映射关系却已被记录，说明不匹配
			for _, v := range mp {
				if v == pattern[i] {
					return false
				}
			}
			mp[v] = pattern[i]
		}
	}
	return true
}
