// k0 [0]         +[1]          =[1,1,1,1,1]
// k1 [1,1,1,1,1] +[0,1,2,3,4]  =[1,2,3,4,5]
// k2 [1,2,3,4,5] +[0,1,3,6,10]
// k3 []+[0,1,4,10,20]
// k4 +[0,1,5,15,35]

class Solution {
public:
    int mod= 1000000007;
    int valueAfterKSeconds(int n, int k) {
        vector<vector<int>> mat(k+1, vector<int>(n,1));

        vector<int> arr(n, 1);
        for (int sec = 1; sec <= k; sec++) {
          int count = 1;
          for (int i = 1; i < n; i++) {
            count = (mat[sec-1][i] + count) % mod;

            mat[sec][i] = count;
          }
        } 

        return mat[k][n-1];
    }
};
