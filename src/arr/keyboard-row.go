package arr

import "unicode"

var mp = map[rune]int{
	'q': 1,
	'w': 1,
	'e': 1,
	'r': 1,
	't': 1,
	'y': 1,
	'u': 1,
	'i': 1,
	'o': 1,
	'p': 1,
	'a': 2,
	's': 2,
	'd': 2,
	'f': 2,
	'g': 2,
	'h': 2,
	'j': 2,
	'k': 2,
	'l': 2,
	'z': 3,
	'x': 3,
	'c': 3,
	'v': 3,
	'b': 3,
	'n': 3,
	'm': 3,
}

func FindWords(words []string) []string {
	var resp []string
	for _, v := range words {
		if check(v) {
			resp = append(resp, v)
		}
	}
	return resp
}

func check(str string) bool {
	h := rune(str[0])
	if unicode.IsUpper(h) {
		h = unicode.ToLower(h)
	}
	head := mp[h]
	for i := 1; i < len(str); i++ {
		cur := rune(str[i])
		if unicode.IsUpper(cur) {
			cur = unicode.ToLower(cur)
		}
		if mp[cur] != head {
			return false
		}
	}
	return true
}
