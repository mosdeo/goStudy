package myLeetCode

func twoSum(nums []int, target int) []int {
	var out []int
	var lengthOfNums = len(nums)

	for i := 0; i < lengthOfNums; i++ {
		for delta := 1; i+delta < lengthOfNums; delta++ {
			if target == nums[i]+nums[i+delta] {
				out = []int{i, i + delta}
				i = lengthOfNums
				break

			}
		}
	}

	return out
}
