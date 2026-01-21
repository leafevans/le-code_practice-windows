/*
 * @lc app=leetcode.cn id=3 lang=cpp
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
#include <algorithm>
#include <cstring>
#include <string>
#include <unordered_map>
using namespace std;
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int last[256];
        fill_n(begin(last), 256, -1);
        int maxLen = 0;
        int left = 0;
        for (int right = 0; right < s.size(); ++right) {
            if (last[s[right]] >= left) {
                left = last[s[right]] + 1;
            }
            last[s[right]] = right;
            maxLen = max(maxLen, right - left + 1);
        }
        return maxLen;
    }
};
// @lc code=end
