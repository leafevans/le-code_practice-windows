/*
 * @lc app=leetcode.cn id=70 lang=cpp
 *
 * [70] 爬楼梯
 */

// @lc code=start
class Solution {
public:
    int climbStairs(int n) {
        memo.assign(n + 1, -1);
        return dfs(n);
    }

private:
    vector<int> memo;

    int dfs(int k) {
        if (k <= 2) {
            return k;
        }
        if (memo[k] != -1) {
            return memo[k];
        }
        return memo[k] = dfs(k - 1) + dfs(k - 2);
    }
};
// @lc code=end
