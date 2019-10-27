package main

import "github.com/aggaer/leetcode/src/arr"

func main() {
	for _, e := range arr.Intersect([]int{4,9,5}, []int{9,4,9,8,4}) {
		println(e)
	}
}
