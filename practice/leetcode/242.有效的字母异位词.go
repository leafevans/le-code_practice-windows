/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [26]int{}
	for _, char := range s {
		count[char-'a']++
	}
	for _, char := range t {
		count[char-'a']--
		if count[char-'a'] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end

