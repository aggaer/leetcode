package main

import "arr"

func main() {
	target := []int{0, 1, 0, 3, 12}
	arr.MoveZeroes(target)
	for _, v := range target {
		println(v)
	}
}
