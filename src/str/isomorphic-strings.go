package str

import "strings"

func IsIsomorphic(s string, t string) bool {
	if len(s) == 1 {
		return true
	}
	for i := 0; i < len(s); i++ {
		if strings.Index(s[:i+1], s[i:i+1]) != strings.Index(t[:i+1], t[i:i+1]) {
			return false
		}
	}
	return true
}
