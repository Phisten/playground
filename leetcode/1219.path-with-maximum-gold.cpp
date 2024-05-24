/*
 * @lc app=leetcode id=1219 lang=cpp
 *
 * [1219] Path with Maximum Gold
 */

// @lc code=start
class Solution {
public:
    int getMaximumGold(vector<vector<int>>& grid) {
        int m = grid.size();
        int n = grid[0].size();

        function<int(int,int,int)> dfs = [&](int i, int j, int score) -> int {

            int tmp = grid[i][j];
            if (tmp == 0) return score;

            grid[i][j] = 0;

            score += tmp;

            int cur = 0;
            if (i - 1 >= 0) 
                cur = max(dfs(i-1,j,score), cur);
            if (i + 1 < m)
                cur = max(dfs(i+1,j,score), cur);
            if (j - 1 >= 0)
                cur = max(dfs(i,j-1,score), cur);
            if (j + 1 < n)
                cur = max(dfs(i,j+1,score), cur);

            grid[i][j] = tmp;

            return cur;
        };

        int ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                ans = max(dfs(i,j,0), ans);
            }
        }
        
        return ans;
    }
};
// @lc code=end

