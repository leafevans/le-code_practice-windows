/*
 * @lc app=leetcode.cn id=39 lang=cpp
 *
 * [39] 组合总和
 */

// @lc code=start
class Solution {
public:
    vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
        backtrack(candidates, target, 0);
        return res;
    }

private:
    vector<int> path;
    vector<vector<int>> res;

    void backtrack(vector<int>& candidates, int target, int start) {
        if (target < 0) {
            return;
        }
        if (target == 0) {
            res.push_back(path);
            return;
        }

        for (int i = start; i < candidates.size(); ++i) {
            path.push_back(candidates[i]);
            backtrack(candidates, target - candidates[i], i);
            path.pop_back();
        }
    }
};
// @lc code=end
