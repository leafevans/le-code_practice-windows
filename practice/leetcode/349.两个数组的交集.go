/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	count1 := [1001]int{}
	count2 := [1001]int{}
	res := make([]int, 0)
	for _, v := range nums1 {
		count1[v] = 1
	}
	for _, v := range nums2 {
		count2[v] = 1
	}
	for i := range 1001 {
		if count1[i]+count2[i] == 2 {
			res = append(res, i)
		}
	}
	return res
}

// @lc code=end

