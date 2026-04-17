/*
 * @lc app=leetcode.cn id=28 lang=cpp
 *
 * [28] 找出字符串中第一个匹配项的下标
 */

// @lc code=start
class Solution {
public:
    int strStr(string haystack, string needle) {
        vector<int> next = getNext(needle);
        int m = haystack.size(), n = needle.size();
        int j = 0;
        for (int i = 0; i < m; ++i) {
            while (j > 0 && haystack[i] != needle[j]) {
                j = next[j - 1];
            }
            if (haystack[i] == needle[j]) {
                ++j;
            }
            if (j == n) {
                return i - n + 1;
            }
        }
        return -1;
    }

private:
    vector<int> getNext(string needle) {
        int n = needle.size();
        vector<int> next(n, 0);
        int j = 0;
        for (int i = 1; i < n; ++i) {
            while (j > 0 && needle[i] != needle[j]) {
                j = next[j - 1];
            }
            if (needle[i] == needle[j]) {
                ++j;
            }
            next[i] = j;
        }
        return
    }
};
// @lc code=end
