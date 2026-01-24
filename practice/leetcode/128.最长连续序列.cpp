/*
 * @lc app=leetcode.cn id=128 lang=cpp
 *
 * [128] 最长连续序列
 */

// @lc code=start
#include <algorithm>
#include <unordered_set>
#include <vector>
using namespace std;
class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        if (nums.empty()) {
            return 0;
        }
        unordered_set<int> s(nums.begin(), nums.end());
        int maxLen = 0;
        for (int num : s) {
            if (!s.count(num - 1)) {
                int curr = num, currLen = 1;
                while (s.count(curr + 1)) {
                    ++curr;
                    ++currLen;
                }
                maxLen = max(maxLen, currLen);
            }
        }
        return maxLen;
    }
};
// @lc code=end
