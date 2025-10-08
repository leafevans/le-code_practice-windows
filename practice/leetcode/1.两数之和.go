/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, exists := hashMap[complement]; exists {
			return []int{i, j}
		}
		hashMap[num] = i
	}

	return []int{}

}

// @lc code=end

