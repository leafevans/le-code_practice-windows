/*
 * @lc app=leetcode.cn id=394 lang=cpp
 *
 * [394] 字符串解码
 */

// @lc code=start
class Solution {
public:
    string decodeString(string s) {
        int num = 0;
        string res = "";
        stack<int> nums;
        stack<string> strs;

        for (char c : s) {
            if (isdigit(c)) {
                num = num * 10 + c - '0';
            } else if (c == '[') {
                nums.push(num);
                strs.push(res);
                num = 0;
                res = "";
            } else if (c == ']') {
                int times = nums.top();
                nums.pop();

                string prev_str = strs.top();
                strs.pop();

                for (int i = 0; i < times; i++) {
                    prev_str += res;
                }

                res = prev_str;
            } else {
                res += c;
            }
        }

        return res;
    }
};
// @lc code=end
