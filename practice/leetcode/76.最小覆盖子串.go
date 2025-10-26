/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	lenS, lenT := len(s), len(t)
	if lenS == 0 || lenT == 0 {
		return ""
	}

	need := make(map[byte]int)
	for i := range lenT {
		need[t[i]]++
	}

	start, left, valid, minLen := 0, 0, 0, lenS+1
	window := make(map[byte]int)
	for right := range lenS {
		c := s[right]
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			subLen := right - left + 1
			if subLen < minLen {
				minLen = subLen
				start = left
			}

			d := s[left]
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
			left++
		}
	}

	if minLen == lenS+1 {
		return ""
	}
	return s[start : start+minLen]
}

// @lc code=end

