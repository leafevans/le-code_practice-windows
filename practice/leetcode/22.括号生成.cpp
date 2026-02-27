/*
 * @lc app=leetcode.cn id=22 lang=cpp
 *
 * [22] 括号生成
 */

// @lc code=start
class Solution {
public:
    vector<string> generateParenthesis(int n) {
        backtrack(n, 0, 0);
        return res;
    }

private:
    vector<string> res;
    string path;

    void backtrack(int n, int open, int close) {
        if (open == n && close == n) {
            res.push_back(path);
            return;
        }

        if (open < n) {
            path.push_back('(');
            backtrack(n, open + 1, close);
            path.pop_back();
        }

        if (close < open) {
            path.push_back(')');
            backtrack(n, open, close + 1);
            path.pop_back();
        }
    }
};
// @lc code=end
