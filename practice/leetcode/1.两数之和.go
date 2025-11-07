/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			return []int{i, j}
		}
		m[num] = i
	}
	return nil
}

// @lc code=end

