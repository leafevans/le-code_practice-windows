/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	rows, cols := len(matrix), len(matrix[0])
	res := make([]int, rows*cols)
	k := 0

	top, bottom := 0, rows-1
	left, right := 0, cols-1

	for top <= bottom && left <= right {
		for j := left; j <= right; j++ {
			res[k] = matrix[top][j]
			k++
		}
		top++

		for i := top; i <= bottom; i++ {
			res[k] = matrix[i][right]
			k++
		}
		right--

		if top <= bottom {
			for j := right; j >= left; j-- {
				res[k] = matrix[bottom][j]
				k++
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res[k] = matrix[i][left]
				k++
			}
			left++
		}
	}

	return res
}

// @lc code=end

