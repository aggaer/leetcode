package main

import "arr"

func main() {
	resp := arr.ReorderLogFiles([]string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo","a2 act car"})
	for _, v := range resp {
		println(v)
	}
}
