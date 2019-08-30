package others

func missingNumber(nums []int) int {
	expect := (len(nums)) * (len(nums) + 1) / 2
	actual := 0
	for _, v := range nums {
		actual += v
	}
	return expect - actual
}
