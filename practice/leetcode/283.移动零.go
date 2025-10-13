/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	k, n := 0, len(nums)
	for i := range n {
		if nums[i] != 0 {
			if i != k {
				nums[i], nums[k] = nums[k], nums[i]
			}
			k++
		}
	}
}

// @lc code=end

