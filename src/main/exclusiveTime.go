package main

import "leetcode/src/str"



func main() {
	req := []string{"0:start:0",
		"1:start:2",
		"1:end:5",
		"0:end:6"}
	resp := str.ExclusiveTime(2, req)
	for _, v := range resp {
		println(v)
	}
}
