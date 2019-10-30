package str

//给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。
// 如果剩余少于 k 个字符，则将剩余的所有全部反转。
// 如果有小于 2k 但大于或等于 k 个字符，则反转前 k 个字符，并将剩余的字符保持原样
func ReverseStr(s string, k int) string {
	bytes := []byte(s)
	for i := 0; i < len(bytes); i += 2 * k {
		if i+k < len(bytes) {
			ReverseString(bytes[i : i+k])
		} else {
			ReverseString(bytes[i:])
		}
	}
	return string(bytes)
}
