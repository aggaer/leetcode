package arr

import "github.com/aggaer/leetcode/src/utils"

//给定一个二进制数组， 计算其中最大连续1的个数。
//输入: [1,1,0,1,1,1]
//输出: 3
//解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
func FindMaxConsecutiveOnes(nums []int) int {
	resp, num := 0, 0
	for _, v := range nums {
		if v == 1 {
			num++
		} else {
			resp = utils.Max(resp, num)
			num = 0
		}
	}
	if resp > num {
		return resp
	} else {
		return num
	}
}
