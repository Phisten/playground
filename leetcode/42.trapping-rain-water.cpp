/*
 * @lc app=leetcode id=42 lang=cpp
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
class Solution {
public:
    int trap(vector<int>& height) {
        int n = height.size();
        vector<int> lh(n, 0), rh(n ,0);
        for (int i = 1; i < n; i++) {
            lh[i] = max(height[i - 1], lh[i - 1]);
            int revI = n - i - 1;
            rh[revI] = max(height[revI + 1], rh[revI + 1]);
        }

        int ans = 0;
        for (int i = 1; i < n - 1; i++) {
            int cur = max(0, min(lh[i], rh[i]) - height[i]);
            ans += cur;
            // cout << "i=" << i << " lh=" <<
            //     lh[i] << " rh=" << rh[i] << " h=" << height[i] << 
            //     ", cur=" << cur << " ans:" << ans << endl;
        }
        
        return ans;
    }
};
// @lc code=end

