package str

//noinspection GoUnusedFunction,SpellCheckingInspection
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	len1 := len(a)
	len2 := len(b)
	if len1 > len2 {
		return len1
	} else {
		return len2
	}
}
