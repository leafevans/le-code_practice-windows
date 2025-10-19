/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	start := 0
	minLen := n + 1
	currentSum := 0
	for end := range n {
		currentSum += nums[end]
		for currentSum >= target {
			subLen := end - start + 1
			if subLen < minLen {
				minLen = subLen
			}
			currentSum -= nums[start]
			start++
		}
	}

	if minLen == n+1 {
		return 0
	}
	return minLen
}

// @lc code=end

