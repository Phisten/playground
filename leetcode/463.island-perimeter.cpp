/*
 * @lc app=leetcode id=463 lang=cpp
 *
 * [463] Island Perimeter
 */

// @lc code=start
class Solution {
public:
    int islandPerimeter(vector<vector<int>>& grid) {
        int n = grid.size();
        int col = grid[0].size();

        int ans = 0;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < col; j++) {
                if (i == n - 1 && grid[i][j] == 1)
                    ans++;
                if (i == 0 && grid[i][j] == 1)
                    ans++;
                if (i > 0 && grid[i][j] != grid[i - 1][j])
                    ans++;
                if (j == 0 && grid[i][j] == 1)
                    ans++;
                if (j == col - 1 && grid[i][j] == 1)
                    ans++;
                if (j > 0 && grid[i][j] != grid[i][j - 1])
                    ans++;
            }
        }
        return ans;
    }
};
// @lc code=end
