/*
 * @lc app=leetcode.cn id=79 lang=cpp
 *
 * [79] 单词搜索
 */

// @lc code=start
class Solution {
public:
    bool exist(vector<vector<char>>& board, string word) {
        int m = board.size(), n = board[0].size();
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (dfs(board, word, 0, i, j)) {
                    return true;
                }
            }
        }
        return false;
    }

private:
    bool dfs(vector<vector<char>>& board, const string& word, int pos, int i, int j) {
        if (pos == word.size()) {
            return true;
        }

        int m = board.size(), n = board[0].size();
        if (i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[pos]) {
            return false;
        }

        char tmp = board[i][j];
        board[i][j] = '#';

        bool found = dfs(board, word, pos + 1, i - 1, j) || dfs(board, word, pos + 1, i + 1, j) ||
                     dfs(board, word, pos + 1, i, j - 1) || dfs(board, word, pos + 1, i, j + 1);

        board[i][j] = tmp;

        return found;
    }
};
// @lc code=end
