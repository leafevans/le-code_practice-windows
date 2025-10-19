/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start

func sortedSquares(nums []int) []int {
	n := len(nums)
	left, right, pos := 0, n-1, n-1
	res := make([]int, n)

	for left <= right {
		leftSq := nums[left] * nums[left]
		rightSq := nums[right] * nums[right]
		if leftSq >= rightSq {
			res[pos] = leftSq
			left++
		} else {
			res[pos] = rightSq
			right--
		}
		pos--
	}

	return res
}

// @lc code=end

