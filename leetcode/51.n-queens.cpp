/*
 * @lc app=leetcode id=51 lang=cpp
 *
 * [51] N-Queens
 */

// @lc code=start
#include <vector>
#include <functional>
using namespace std;

class Solution
{
public:
    vector<vector<string>> solveNQueens(int n)
    {
        vector<vector<string>> result;
        vector<string> board(n, string(n, '.'));

        vector<vector<bool>> slot;
        for (int i = 0; i < n; i++)
            slot.push_back(vector<bool>(n, false));

        auto isValid = [&](int y, int x)
        {
            for (int i = 1; i <= y; i++)
            {
                if (board[y - i][x] == 'Q')
                    return false;
                if (x + i < n && board[y - i][x + i] == 'Q')
                    return false;
                if (x - i >= 0 && board[y - i][x - i] == 'Q')
                    return false;
            }
            return true;
        };

        function<void(int)>
            backtrack = [&](int y)
        {
            if (y == n)
            {
                result.push_back(board);
                return;
            }

            for (int x = 0; x < n; x++)
            {
                if (isValid(y, x))
                {
                    board[y][x] = 'Q';
                    backtrack(y + 1);
                    board[y][x] = '.';
                }
            }
        };

        backtrack(0);
        return result;
    }
};
// @lc code=end
