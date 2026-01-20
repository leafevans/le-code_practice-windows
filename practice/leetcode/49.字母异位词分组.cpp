/*
 * @lc app=leetcode.cn id=49 lang=cpp
 *
 * [49] 字母异位词分组
 */

// @lc code=start
#include <algorithm>
#include <string>
#include <unordered_map>
#include <vector>
using namespace std;
class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        unordered_map<string, vector<string>> mp;
        for (auto& str : strs) {
            string key(26, '0');
            for (char c : str) {
                key[c - 'a']++;
            }
            mp[key].push_back(str);
        }
        vector<vector<string>> res;
        for (auto& p : mp) {
            res.push_back(p.second);
        }
        return res;
    }
};
// @lc code=end
