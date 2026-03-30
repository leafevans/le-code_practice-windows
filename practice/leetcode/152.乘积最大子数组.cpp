/*
 * @lc app=leetcode.cn id=152 lang=cpp
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
class Solution {
public:
    int maxProduct(vector<int>& nums) {
        int res = nums[0];
        int curMax = nums[0];
        int curMin = nums[0];

        for (int i = 1; i < nums.size(); ++i) {
            int x = nums[i];
            int prevMax = curMax;
            int prevMin = curMin;

            curMax = max({x, prevMax * x, prevMin * x});
            curMin = min({x, prevMax * x, prevMin * x});

            res = max(res, curMax);
        }

        return res;
    }
};
// @lc code=end
