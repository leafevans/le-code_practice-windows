/*
 * @lc app=leetcode.cn id=3 lang=cpp
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int n = s.size();
        if (n <= 1) {
            return n;
        }
        int last[256];
        for (int i = 0; i < 256; ++i) {
            last[i] = -1;
        }
        int maxLen = 0;
        int l = 0;
        for (int r = 0; r < n; ++r) {
            if (last[s[r]] >= l) {
                l = last[s[r]] + 1;
            }
            last[s[r]] = r;
            maxLen = max(maxLen, r - l + 1);
        }
        return maxLen;
    }
};
// @lc code=end
