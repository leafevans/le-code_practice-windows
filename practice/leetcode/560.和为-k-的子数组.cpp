/*
 * @lc app=leetcode.cn id=560 lang=cpp
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
#include <unordered_map>
#include <vector>
using namespace std;
class Solution {
public:
    int subarraySum(vector<int>& nums, int k) {
        unordered_map<int, int> preSumMap;
        preSumMap[0] = 1;
        int sum = 0, count = 0;
        for (int num : nums) {
            sum += num;
            if (preSumMap.find(sum - k) != preSumMap.end()) {
                count += preSumMap[sum - k];
            }
            ++preSumMap[sum];
        }
        return count;
    }
};
// @lc code=end
