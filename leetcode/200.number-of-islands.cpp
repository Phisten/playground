/*
 * @lc app=leetcode id=200 lang=cpp
 *
 * [200] Number of Islands
 */

// @lc code=start
class Solution
{
public:
    int numIslands(vector<vector<char>> &grid)
    {
        int ans = 0;
        int m = grid.size();
        int n = grid[0].size();

        function<void(int, int)> filler = [&](int i, int j)
        {
            if (grid[i][j] != '1')
                return;
            grid[i][j] = '0';

            if (i > 0)
                filler(i - 1, j);
            if (i < m - 1)
                filler(i + 1, j);
            if (j > 0)
                filler(i, j - 1);
            if (j < n - 1)
                filler(i, j + 1);
        };

        for (int i = 0; i < m; i++)
        {
            for (int j = 0; j < n; j++)
            {
                if (grid[i][j] == '1')
                {
                    ans++;
                    filler(i, j);
                }
            }
        }

        return ans;
    }
};
// @lc code=end
