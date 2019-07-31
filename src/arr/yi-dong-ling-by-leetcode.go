package arr

//慢指针之前所有元素都是非零的
//快指针和慢指针之间所有元素都是0
func MoveZeroes(nums []int) {
	for slow, fast := 0, 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
