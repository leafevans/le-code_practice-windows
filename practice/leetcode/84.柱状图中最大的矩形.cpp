/*
 * @lc app=leetcode.cn id=84 lang=cpp
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
class Solution {
public:
    int largestRectangleArea(vector<int>& heights) {
        int n = heights.size();
        vector<int> h(n + 2, 0);
        for (int i = 0; i < n; ++i) {
            h[i + 1] = heights[i];
        }
        int ans = 0;
        stack<int> st;
        st.push(0);
        for (int i = 1; i < n + 2; ++i) {
            while (!st.empty() && h[i] < h[st.top()]) {
                int height = h[st.top()];
                st.pop();
                int width = i - st.top() - 1;
                ans = max(ans, height * width);
            }
            st.push(i);
        }
        return ans;
    }
};
// @lc code=end
