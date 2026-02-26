/*
 * @lc app=leetcode.cn id=46 lang=cpp
 *
 * [46] 全排列
 */

// @lc code=start
class Solution {
public:
    vector<vector<int>> permute(vector<int>& nums) {
        backtrack(nums, 0);
        return res;
    }

private:
    vector<vector<int>> res;

    void backtrack(vector<int>& nums, int start) {
        if (start == nums.size()) {
            res.push_back(nums);
            return;
        }

        for (int i = start; i < nums.size(); ++i) {
            swap(nums[i], nums[start]);
            backtrack(nums, start + 1);
            swap(nums[i], nums[start]);
        }
    }
};
// @lc code=end
