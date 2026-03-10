/*
 * @lc app=leetcode.cn id=51 lang=cpp
 *
 * [51] N 皇后
 */

// @lc code=start
class Solution {
public:
    vector<vector<string>> solveNQueens(int n) {
        vector<string> board(n, string(n, '.'));
        backtrack(n, 0, board);
        return res;
    }

private:
    vector<vector<string>> res;

    void backtrack(int n, int row, vector<string>& board) {
        if (row == n) {
            res.push_back(board);
            return;
        }

        for (int col = 0; col < n; ++col) {
            if (isValid(board, row, col, n)) {
                board[row][col] = 'Q';
                backtrack(n, row + 1, board);
                board[row][col] = '.';
            }
        }
    }

    bool isValid(const vector<string>& board, int row, int col, int n) {
        for (int i = 0; i < row; ++i) {
            if (board[i][col] == 'Q') {
                return false;
            }
        }

        for (int i = row - 1, j = col - 1; i >= 0 && j >= 0; --i, --j) {
            if (board[i][j] == 'Q') {
                return false;
            }
        }

        for (int i = row - 1, j = col + 1; i >= 0 && j < n; --i, ++j) {
            if (board[i][j] == 'Q') {
                return false;
            }
        }

        return true;
    }
};
// @lc code=end
