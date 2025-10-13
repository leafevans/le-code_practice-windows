/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	slow, n := 0, len(nums)
	for fast := 1; fast < n; fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

// @lc code=end

