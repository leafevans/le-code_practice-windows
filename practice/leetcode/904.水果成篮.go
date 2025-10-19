/*
 * @lc app=leetcode.cn id=904 lang=golang
 *
 * [904] 水果成篮
 */

// @lc code=start
func totalFruit(fruits []int) int {
	n := len(fruits)
	if n == 0 {
		return 0
	}

	fruitCount := make(map[int]int)
	start, maxLen := 0, 0
	for end := range n {
		fruitCount[fruits[end]]++

		for len(fruitCount) > 2 {
			fruitCount[fruits[start]]--
			if fruitCount[fruits[start]] == 0 {
				delete(fruitCount, fruits[start])
			}
			start++
		}

		subLen := end - start + 1
		if maxLen < subLen {
			maxLen = subLen
		}
	}

	return maxLen
}

// @lc code=end

