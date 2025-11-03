/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	set := make(map[int]struct{})
	for n != 1 {
		if _, ok := set[n]; ok {
			return false
		}
		set[n] = struct{}{}
		n = getSum(n)
	}
	return true
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

// @lc code=end

