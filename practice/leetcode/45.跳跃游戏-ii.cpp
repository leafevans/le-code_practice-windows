/*
 * @lc app=leetcode.cn id=45 lang=cpp
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
class Solution {
public:
    int jump(vector<int>& nums) {
        int n = nums.size();
        if (n <= 1) {
            return 0;
        }
        int jumps = 0;
        int end = 0;
        int farthest = 0;
        for (int i = 0; i < n - 1; ++i) {
            farthest = max(farthest, i + nums[i]);
            if (i == end) {
                ++jumps;
                end = farthest;
                if (end >= n - 1) {
                    break;
                }
            }
        }
        return jumps;
    }
};
// @lc code=end
