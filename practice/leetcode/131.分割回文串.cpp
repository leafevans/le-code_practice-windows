/*
 * @lc app=leetcode.cn id=131 lang=cpp
 *
 * [131] 分割回文串
 */

// @lc code=start
class Solution {
public:
    vector<vector<string>> partition(string s) {
        int n = s.size();
        vector<vector<bool>> isPal(n, vector<bool>(n, false));
        for (int i = n - 1; i >= 0; --i) {
            for (int j = i; j < n; ++j) {
                if (s[i] == s[j] && (j - i < 2 || isPal[i + 1][j - 1])) {
                    isPal[i][j] = true;
                }
            }
        }

        dfs(s, 0, isPal);
        return res;
    }

private:
    vector<vector<string>> res;
    vector<string> path;

    void dfs(const string& s, int start, const vector<vector<bool>>& isPal) {
        int n = s.size();
        if (start == n) {
            res.push_back(path);
            return;
        }

        for (int end = start; end < n; ++end) {
            if (isPal[start][end]) {
                path.push_back(s.substr(start, end - start + 1));
                dfs(s, end + 1, isPal);
                path.pop_back();
            }
        }
    }
};
// @lc code=end
