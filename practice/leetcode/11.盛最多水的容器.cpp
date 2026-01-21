/*
 * @lc app=leetcode.cn id=11 lang=cpp
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
#include <algorithm>
#include <vector>
using namespace std;
class Solution {
public:
    int maxArea(vector<int>& height) {
        int left = 0, right = height.size() - 1;
        int maxArea = 0;
        while (left < right) {
            int h = min(height[left], height[right]);
            int w = right - left;
            maxArea = max(maxArea, h * w);
            if (height[left] < height[right]) {
                ++left;
            } else {
                --right;
            }
        }
        return maxArea;
    }
};
// @lc code=end
