/*
 * @lc app=leetcode id=37 lang=cpp
 *
 * [37] Sudoku Solver
 */

// @lc code=start
class Solution {
public:
    void solveSudoku(vector<vector<char>>& board) {
        function<bool(int, int)> isValid = [&](int y, int x) {
            // check y and x and 3x3
            int xBase = x - x % 3;
            int yBase = y - y % 3; 
            for (int i = 0; i < 9; i++) {
                if (board[i][x] == board[y][x] and i != y) return false;
                if (board[y][i] == board[y][x] and i != x) return false;
                if (board[yBase + i/3][xBase + i%3] == board[y][x] and (yBase + i/3) != y and (xBase + i%3) != x) return false;
            }
            return true;
        };

        function<bool(int, int)> fillNum = [&](int y, int x) {
            if (y == 9) return true;
            if (x == 9) return fillNum(y+1, 0);

            if (board[y][x] == '.') {
                for (int k = 1; k <= 9; k++) {
                    board[y][x] = '0' + k;
                    if (isValid(y, x) and fillNum(y, x+1)) {
                        return true;
                    }
                }
                board[y][x] = '.';
                return false;
            } else {
                return isValid(y, x) and fillNum(y, x+1);
            }
        };

        fillNum(0, 0);
    }
};
// @lc code=end
