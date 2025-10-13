/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left, right := 1, x
	for left < right {
		mid := left + (right-left)>>1
		if mid <= x/mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}

// @lc code=end

