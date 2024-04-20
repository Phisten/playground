/*
 * @lc app=leetcode id=1992 lang=cpp
 *
 * [1992] Find All Groups of Farmland
 */

// @lc code=start
class Solution {
public:
    vector<vector<int>> findFarmland(vector<vector<int>>& land) {
        int m = land.size();
        int n = land[0].size();

        int count = 0;
        vector<int> pos(4);
        function<void(int,int,int)> blob = [&](int i, int j, int label) {
            if (land[i][j] != 1) return;
            land[i][j] = label;

            if (i < pos[0]) pos[0] = i;
            if (i > pos[2]) pos[2] = i;
            if (j < pos[1]) pos[1] = j;
            if (j > pos[3]) pos[3] = j;

            if (i > 0) blob(i - 1, j, label);
            if (i < m - 1) blob(i + 1, j, label);
            if (j > 0) blob(i, j - 1, label);
            if (j < n - 1) blob(i, j + 1, label);
        };

        vector<vector<int>> ans;
        for (int i = 0; i < m ; i++) {
            for (int j = 0; j < n; j++) {
                if (land[i][j] == 1) {
                    count++;
                    pos[0] = i;
                    pos[1] = j;
                    pos[2] = i;
                    pos[3] = j;
                    blob(i, j, 10 + count);
                    ans.push_back(pos);
                }
            }
        }

        return ans;
    }
};
// @lc code=end

