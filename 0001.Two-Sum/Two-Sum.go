package leetcode

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			first := nums[i]
			next := nums[j]
			sum := first+next
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
