class Solution {
public:
    
    bool canMakeSquare(vector<vector<char>>& grid) {
        int n = 3;
        
        for (int i = 0; i < 2; i++) {
            for (int j = 0; j < 2; j++) {
                int w = 0;
                int b = 0;
                if (grid[i][j] == 'W') w++;
                else b++;
                if (grid[i+1][j] == 'W') w++;
                else b++;
                if (grid[i][j+1] == 'W') w++;
                else b++;
                if (grid[i+1][j+1] == 'W') w++;
                else b++;
                
                if (w >= 3 || b >= 3) return true;
                
            }
        }
        return false;
    }
};
