class Solution {
public:
    int maxHeightOfTriangle(int red, int blue) {
        
        int n = 1;
        int a = 0;
        int b = 0;
        int minBall = min(red, blue);
        int maxBall = max(red, blue);
        int ans = 0;
        
        while (true) {
            if (n % 2 == 1) {
                a += n;
            } else {
                b += n;
            }
            
            if (a <= minBall && b <= maxBall) ans++;
            else if (b <= minBall && a <= maxBall) ans++;
            else break;
            
            n++;
        }
        
        return ans;
    }
};
