/*
 * @lc app=leetcode.cn id=5 lang=cpp
 *
 * [5] 最长回文子串
 */

// @lc code=start
class Solution {
public:
    string longestPalindrome(string s) {
        int n = s.size();
        if (n <= 1) {
            return s;
        }

        int baseL = 0, baseR = 0;

        auto expand = [&](int l, int r) {
            while (l >= 0 && r < n && s[l] == s[r]) {
                --l;
                ++r;
            }
            int len = r - l - 1;
            if (len > baseR - baseL) {
                baseL = l + 1;
                baseR = r;
            }
        };

        for (int i = 0; i < n; ++i) {
            expand(i, i);
            expand(i, i + 1);
        }

        return s.substr(baseL, baseR - baseL);
    }
};
// @lc code=end
