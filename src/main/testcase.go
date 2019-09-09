package main

import "leetcode/src/arr"

//"---"
//nums1 = [4,1,2], nums2 = [1,3,4,2].
//3
func main() {
	//println(str.LicenseKeyFormatting("---", 3))
	//println(arr.FindMaxConsecutiveOnes([]int{1,0,1,1}))
	//println(others.ConstructRectangle(12450)[1])
	resp := arr.NextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})
	for _, v := range resp {
		println(v)
	}
}
