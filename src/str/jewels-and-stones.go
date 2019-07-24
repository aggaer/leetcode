package str

import "strings"

func NumJewelsInStones(J string, S string) int {
	sum := 0
	for _, v := range S {
		if strings.ContainsRune(J, v) {
			sum++
		}
	}
	return sum
}
