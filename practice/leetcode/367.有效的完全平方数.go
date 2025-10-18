/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	left, right := 0, num
	for left < right {
		mid := left + (right-left)>>1
		if mid <= num/mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return num == (left-1)*(left-1)
}

// @lc code=end

