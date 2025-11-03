/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			m[a+b]++
		}
	}
	count := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			if val, ok := m[-(c + d)]; ok {
				count += val
			}
		}
	}
	return count
}

// @lc code=end

