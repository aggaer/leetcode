package str

func TitleToNumber(s string) int {
	resp := 0
	for _, v := range s {
		num := v - 'A' + 1
		resp = resp*26 + int(num)
	}
	return resp
}
