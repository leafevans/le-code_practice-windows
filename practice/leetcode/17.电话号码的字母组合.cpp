/*
 * @lc app=leetcode.cn id=17 lang=cpp
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
class Solution {
public:
    vector<string> letterCombinations(string digits) {
        backtrack(digits, 0);
        return res;
    }

private:
    vector<string> res;
    string path;
    const vector<string> MAPPING = {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};

    void backtrack(string digits, int idx) {
        if (idx == digits.size()) {
            res.push_back(path);
            return;
        }

        int digit = digits[idx] - '0';
        const string letters = MAPPING[digit];

        for (char ch : letters) {
            path.push_back(ch);
            backtrack(digits, idx + 1);
            path.pop_back();
        }
    }
};
// @lc code=end
