/*
 * @lc app=leetcode id=1289 lang=cpp
 *
 * [1289] Minimum Falling Path Sum II
 */

// @lc code=start
class Solution {
public:
    int minFallingPathSum(vector<vector<int>>& grid) {
        int n = grid.size();

        vector<int> lastRow(n);
        vector<int> curRow(n);
        for (int i = 0; i < n; i++) {
            lastRow[i] = grid[0][i];
        }

        for (int i = 1; i < n; i++) {
            for (int j = 0; j < n; j++) {
                int minLast = INT_MAX;
                for (int k = 0; k < n; k++)
                    if (k != j)
                        minLast = min(minLast, lastRow[k]);

                curRow[j] = grid[i][j] + minLast;
            }

            for (int j = 0; j < n; j++)
                lastRow[j] = curRow[j];
        }

        int ans = lastRow[0];
        for (int i = 1; i < n; i++)
            ans = min(lastRow[i], ans);

        return ans;
    }
};
// @lc code=end

