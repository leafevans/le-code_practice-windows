/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
// ...existing code...
func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	skipS, skipT := 0, 0

	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				i--
				skipS++
			} else if skipS > 0 {
				i--
				skipS--
			} else {
				break
			}
		}

		for j >= 0 {
			if t[j] == '#' {
				j--
				skipT++
			} else if skipT > 0 {
				j--
				skipT--
			} else {
				break
			}
		}

		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}

	return true
}

// @lc code=end

