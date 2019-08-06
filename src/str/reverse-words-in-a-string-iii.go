package str

func ReverseWords(s string) string {
	resp := ""
	pre := []byte("")
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			pre = append(pre[0:], s[i])
			resp = string(pre) + resp
			pre = pre[len(pre):]
			println("pre: " + string(pre))
		} else {
			pre = append(pre, s[i])
		}
	}
	resp = string(pre) +" "+ resp
	return resp[:len(resp)-1]
}
