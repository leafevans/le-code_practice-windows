/*
 * @lc app=leetcode.cn id=74 lang=cpp
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        int m = matrix.size(), n = matrix[0].size();
        int l = 0, r = m * n - 1;

        while (l <= r) {
            int mid = l + ((r - l) >> 1);

            if (matrix[mid / n][mid % n] == target) {
                return true;
            } else if (matrix[mid / n][mid % n] < target) {
                l = mid + 1;
            } else {
                r = mid - 1;
            }
        }

        return false;
    }
};
// @lc code=end
