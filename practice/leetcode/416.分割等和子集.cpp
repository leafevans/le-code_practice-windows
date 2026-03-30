/*
 * @lc app=leetcode.cn id=416 lang=cpp
 *
 * [416] 分割等和子集
 */

// @lc code=start
class Solution {
public:
    bool canPartition(vector<int>& nums) {
        int total = 0;
        for (int x : nums) {
            total += x;
        }
        if (total % 2 != 0) {
            return false;
        }
        int target = total / 2;

        vector<bool> dp(target + 1, false);
        dp[0] = true;

        for (int x : nums) {
            for (int j = target; j >= x; --j) {
                if (dp[j - x]) {
                    dp[j] = true;
                }
            }
            if (dp[target]) {
                return true;
            }
        }

        return dp[target];
    }
};
// @lc code=end
