/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left >= len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	start := left

	left, right = 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left == 0 || nums[left-1] != target {
		return []int{-1, -1}
	}
	end := left - 1

	return []int{start, end}
}

// @lc code=end

