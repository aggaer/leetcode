package main

import "arr"

func main() {
	target := []int{1, 2, 3, 4, 5}
	arr.Rotate(target, 7)
	for _, v := range target {
		println(v)
	}
}
