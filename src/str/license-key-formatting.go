package str

import (
	"unicode"
)

//给定一个密钥字符串S，只包含字母，数字以及 '-'（破折号）。N 个 '-' 将字符串分成了 N+1 组。给定一个数字 K，重新格式化字符串，
// 除了第一个分组以外，每个分组要包含 K 个字符，第一个分组至少要包含 1 个字符。
// 两个分组之间用 '-'（破折号）隔开，并且将所有的小写字母转换为大写字母。
func LicenseKeyFormatting(S string, K int) string {
	count := 0
	resp := []byte("")
	for i := len(S) - 1; i >= 0; i-- {
		cur := S[i]
		if cur == '-' {
			continue
		}
		r := rune(cur)
		if unicode.IsLetter(r) {
			r := unicode.ToUpper(r)
			cur = uint8(r)
		}
		resp = append(resp, cur)
		count++
		if count == K {
			resp = append(resp, '-')
			count = 0
		}
	}
	ReverseString(resp)
	ret := string(resp)
	if len(ret) != 0 && ret[0] == '-' {
		ret = ret[1:]
	}
	return ret
}
