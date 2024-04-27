class Solution {
public:
    long long numberOfRightTriangles(vector<vector<int>>& grid) {
        long long ans = 0;
        
        int m = grid.size();
        int n = grid[0].size();
        
        vector<int> y_left(m);
        vector<int> y_right(m);
        vector<int> x_top(n);
        vector<int> x_btm(n);

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 1) {
                    y_right[i]++;
                    x_btm[j]++;
                }
            }
        }
        
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 1) {
                    y_right[i]--;
                    x_btm[j]--;
                    
                    ans += y_right[i] * x_btm[j];
                    ans += y_left[i] * x_btm[j];
                    ans += y_right[i] * x_top[j];
                    ans += y_left[i] * x_top[j];
                    
                    y_left[i]++;
                    x_top[j]++;
                }
            }
        }
        return ans;
    }
};
