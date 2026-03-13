/*
 * @lc app=leetcode.cn id=4 lang=cpp
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        if (nums1.size() > nums2.size()) {
            return findMedianSortedArrays(nums2, nums1);
        }

        int m = nums1.size();
        int n = nums2.size();
        int half = (m + n + 1) / 2;

        int left = 0, right = m;
        while (left <= right) {
            int i = left + ((right - left) >> 1);
            int j = half - i;

            int nums1LeftMax = (i == 0 ? INT_MIN : nums1[i - 1]);
            int nums1RightMin = (i == m ? INT_MAX : nums1[i]);
            int nums2LeftMax = (j == 0 ? INT_MIN : nums2[j - 1]);
            int nums2RightMin = (j == n ? INT_MAX : nums2[j]);

            if (nums1LeftMax <= nums2RightMin && nums2LeftMax <= nums1RightMin) {
                int leftMax = max(nums1LeftMax, nums2LeftMax);
                if ((m + n) % 2 == 1) {
                    return leftMax;
                }
                int rightMin = min(nums1RightMin, nums2RightMin);
                return (leftMax + rightMin) / 2.0;
            } else if (nums1LeftMax > nums2RightMin) {
                right = i - 1;
            } else {
                left = i + 1;
            }
        }

        return 0.0;
    }
};
// @lc code=end
