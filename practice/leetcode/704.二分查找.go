/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

// @lc code=end

