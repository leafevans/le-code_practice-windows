/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range n {
		matrix[i] = make([]int, n)
	}

	count := 1
	for layer := 0; layer < n/2; layer++ {
		for j := layer; j < n-1-layer; j++ {
			matrix[layer][j] = count
			count++
		}

		for i := layer; i < n-1-layer; i++ {
			matrix[i][n-1-layer] = count
			count++
		}

		for j := n - 1 - layer; j > layer; j-- {
			matrix[n-1-layer][j] = count
			count++
		}

		for i := n - 1 - layer; i > layer; i-- {
			matrix[i][layer] = count
			count++
		}
	}

	if n%2 == 1 {
		matrix[n/2][n/2] = count
	}

	return matrix
}

// @lc code=end

