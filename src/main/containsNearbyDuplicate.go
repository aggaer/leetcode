package main

import "arr"

func main() {
	ints := []int{1, 2, 3, 1, 2, 3}
	println(arr.ContainsNearbyDuplicate(ints, 2))
}
