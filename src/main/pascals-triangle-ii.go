package main

import "others"

func main() {
	target := others.GetRow(3)
	for _, v := range target {
		println(v)
	}
}
