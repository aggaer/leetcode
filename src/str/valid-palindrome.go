package str

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	rs := []rune(s)
	l := 0
	r := len(s) - 1
	for l < r {
		L := string(rs[l])
		R := string(rs[r])
		println(L)
		println(R)
		if !unicode.IsLetter(rs[l]) &&!unicode.IsDigit(rs[l]) {
			l++
		} else if !unicode.IsLetter(rs[r]) &&!unicode.IsDigit(rs[r]) {
			r--
		} else {
			if !strings.EqualFold(L, R) {
				return false
			}
			l++
			r--
		}
	}
	return true
}
