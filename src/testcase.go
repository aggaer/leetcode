package main

import "github.com/aggaer/leetcode/src/arr"

//"---"
//3
//noinspection SpellCheckingInspection
func main() {
	ranks := arr.FindRelativeRanks([]int{10,3,8,9,4})
	for _, v := range ranks {
		println(v)
	}
}
