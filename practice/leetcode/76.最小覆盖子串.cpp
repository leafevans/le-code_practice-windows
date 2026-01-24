/*
 * @lc app=leetcode.cn id=76 lang=cpp
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
#include <string>
#include <vector>
using namespace std;
class Solution {
public:
    string minWindow(string s, string t) {
        if (s.empty() || t.empty()) {
            return "";
        }
        vector<int> need(128, 0);
        vector<int> wnd(128, 0);
        int count = 0;
        for (char c : t) {
            if (need[c] == 0) {
                ++count;
            }
            ++need[c];
        }
        int left = 0, right = 0;
        int valid = 0;
        int start = 0, len = INT_MAX;
        while (right < s.size()) {
            char c = s[right];
            ++right;
            if (need[c] > 0) {
                ++wnd[c];
                if (wnd[c] == need[c]) {
                    ++valid;
                }
            }
            while (valid == count) {
                if (right - left < len) {
                    len = right - left;
                    start = left;
                }
                char d = s[left];
                ++left;
                if (need[d] > 0) {
                    if (wnd[d] == need[d]) {
                        --valid;
                    }
                    --wnd[d];
                }
            }
        }
        return len == INT_MAX ? "" : s.substr(start, len);
    }
};
// @lc code=end
