/*
 * @lc app=leetcode.cn id=438 lang=cpp
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
class Solution {
public:
    vector<int> findAnagrams(string s, string p) {
        int m = s.size(), n = p.size();
        if (m < n) {
            return {};
        }
        vector<int> res;
        vector<int> sc(26, 0), pc(26, 0);
        for (char c : p) {
            ++pc[c - 'a'];
        }
        for (int i = 0; i < m; ++i) {
            ++sc[s[i] - 'a'];
            if (i >= n) {
                --sc[s[i - n] - 'a'];
            }
            if (sc == pc) {
                res.push_back(i - n + 1);
            }
        }
        return res;
    }
};
// @lc code=end
