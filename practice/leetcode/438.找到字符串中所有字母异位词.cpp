/*
 * @lc app=leetcode.cn id=438 lang=cpp
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
#include <string>
#include <vector>
using namespace std;
class Solution {
public:
    vector<int> findAnagrams(string s, string p) {
        if (s.size() < p.size()) {
            return {};
        }
        vector<int> res;
        vector<int> pc(26, 0), sc(26, 0);
        for (char c : p) {
            ++pc[c - 'a'];
        }
        int wndSize = p.size();
        for (int i = 0; i < s.size(); ++i) {
            ++sc[s[i] - 'a'];
            if (i >= wndSize) {
                --sc[s[i - wndSize] - 'a'];
            }
            if (sc == pc) {
                res.push_back(i - wndSize + 1);
            }
        }
        return res;
    }
};
// @lc code=end
