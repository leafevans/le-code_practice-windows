/*
 * @lc app=leetcode.cn id=32 lang=cpp
 *
 * [32] 最长有效括号
 */

// @lc code=start
class Solution {
public:
    int longestValidParentheses(string s) {
        int n = s.size();
        int maxlen = 0;
        stack<int> st;
        st.push(-1);
        for (int i = 0; i < n; ++i) {
            if (s[i] == '(') {
                st.push(i);
            } else {
                st.pop();
                if (st.empty()) {
                    st.push(i);
                } else {
                    maxlen = max(maxlen, i - st.top());
                }
            }
        }
        return maxlen;
    }
};
// @lc code=end
