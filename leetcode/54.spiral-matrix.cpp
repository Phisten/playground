/*
 * @lc app=leetcode id=54 lang=cpp
 *
 * [54] Spiral Matrix
 */

// @lc code=start
class Solution {
public:
    vector<int> spiralOrder(vector<vector<int>>& matrix) {
        int m = matrix.size();
        int n = matrix[0].size();
        // r b l t
        vector<int> bound(4);
        bound[0] = n - 1;
        bound[1] = m - 1;
        bound[2] = 0;
        bound[3] = 0;

        vector<int> ans;
        int direct = 0;
        int i = 0, j = -1;
        while (true) {
            if (direct == 0) j++;
            else if (direct == 1) i++;
            else if (direct == 2) j--;
            else if (direct == 3) i--;

            if (i == m || i < 0 || j == n || j < 0) break;

            int cur = matrix[i][j];
            if (cur == -111) break;
            matrix[i][j] = -111;
            ans.push_back(cur);

            if (direct == 0) {
                if (j == bound[0]) {
                    bound[3]++;
                    direct = 1;
                }
            }
            else if (direct == 1) {
                if (i == bound[1]) {
                    bound[0]--;
                    direct = 2;
                }
            }
            else if (direct == 2) {
                if (j == bound[2]) {
                    bound[1]--;
                    direct = 3;
                }
            }
            else if (direct == 3) {
                if (i == bound[3]) {
                    bound[2]++;
                    direct = 0;
                }
            }
        }

        return ans;
    }
};
// @lc code=end

