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
	currSum := 0
	minLen := n + 1
	for end := range n {
		currSum += nums[end]
		for currSum >= target {
			subLen := end - start + 1
			if subLen < minLen {
				minLen = subLen
			}
			currSum -= nums[start]
			start++
		}
	}

	if minLen == n+1 {
		return 0
	}
	return minLen
}

// @lc code=end

