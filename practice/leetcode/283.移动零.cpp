/*
 * @lc app=leetcode.cn id=283 lang=cpp
 *
 * [283] 移动零
 */

// @lc code=start
class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        int n = nums.size();
        if (n <= 1) {
            return;
        }
        int j = 0;
        for (int i = 0; i < n; ++i) {
            if (nums[i] != 0) {
                if (i != j) {
                    swap(nums[i], nums[j]);
                }
                ++j;
            }
        }
    }
};
// @lc code=end
