/*
 * @lc app=leetcode id=7 lang=cpp
 *
 * [7] Reverse Integer
 */

// @lc code=start
class Solution {
public:
    int reverse(int x) {
        int maxN = 2147483647 / 10;
        int minN = -2147483648 / 10;
        int lastMod = 7;
        if (x < 0) lastMod = -8;
        
        int ans = 0;
        vector<int> arr;

        while (x > 0 || x < 0) {
            arr.push_back(x % 10);
            x /= 10;
        }

        int n = arr.size();
        for (int i = 0; i < n; i++) {
            int cur = arr[i];
            if ((ans == maxN and cur > lastMod) or (ans == minN and cur < lastMod)) {
                ans = 0;
                break;
            }

            ans += cur;
            if (i + 1 < n) {
                if (ans > maxN || ans < minN) {
                    ans = 0;
                    break;
                }
                ans *= 10;
            }
        }
        
        return ans;
    }
};
// @lc code=end

