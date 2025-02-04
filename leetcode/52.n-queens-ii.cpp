/*
 * @lc app=leetcode id=52 lang=cpp
 *
 * [52] N-Queens II
 */

// @lc code=start
class Solution
{
    public:
        int totalNQueens(int n)
        {
            vector<vector<bool>> board(n, vector<bool>(n));

            function < bool(int, int) > validPos =[&](int y, int x)
            {
                for (int i = 1; i <= y; i++)
                {
                    if (board[y-i][x]) return false;
                    if (x-i >= 0 && board[y-i][x-i]) return false;
                    if (x+i < n && board[y-i][x+i]) return false;
                }
                return true;
            };

            int ans = 0;
            function < void(int) > dsf =[&](int col) {
                if (col == n) {
                    ans++;
                    return;
                }
                for (int x = 0; x < n; x++)
                {
                    if (validPos(col, x)) {
                        board[col][x] = true;
                        dsf(col+1);
                        board[col][x] = false;
                    }
                }
            };
            dsf(0);

            return ans;
        }
};
// @lc code=end
