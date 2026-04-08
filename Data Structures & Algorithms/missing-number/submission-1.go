func missingNumber(nums []int) int {
	l := len(nums)

	sum := 0
	sum = l * (l + 1) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}