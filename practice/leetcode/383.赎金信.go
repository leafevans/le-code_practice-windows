/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	count := [26]int{}
	for _, char := range magazine {
		count[char-'a']++
	}
	for _, char := range ransomNote {
		count[char-'a']--
		if count[char-'a'] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end

