/*
 * @lc app=leetcode.cn id=131 lang=cpp
 *
 * [131] 分割回文串
 */

// @lc code=start
class Solution {
public:
    vector<vector<string>> partition(string s) {
        backtrack(s, 0);
        return res;
    }

private:
    vector<string> path;
    vector<vector<string>> res;

    void backtrack(const string& s, int start) {
        if (start == s.size()) {
            res.push_back(path);
            return;
        }

        for (int i = start; i < s.size(); ++i) {
            if (isPalin(s, start, i)) {
                string str = s.substr(start, i - start + 1);
                path.push_back(str);
                backtrack(s, i + 1);
                path.pop_back();
            }
        }
    }

    bool isPalin(const string& s, int left, int right) {
        while (left < right) {
            if (s[left++] != s[right--]) {
                return false;
            }
        }
        return true;
    }
};
// @lc code=end
