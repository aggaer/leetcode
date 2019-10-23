package str

import "unicode"

func DetectCapitalUse(word string) bool {
	lowerCnt := 0
	for _, v := range word {
		if unicode.IsLower(v) {
			lowerCnt++
		}
	}
	if lowerCnt == len(word) || lowerCnt == 0 {
		return true
	} else {
		return lowerCnt == len(word)-1 && unicode.IsUpper(rune(word[0]))
	}
}
