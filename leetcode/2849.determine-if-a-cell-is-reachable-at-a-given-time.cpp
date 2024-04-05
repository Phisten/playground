/*
 * @lc app=leetcode id=2849 lang=cpp
 *
 * [2849] Determine if a Cell Is Reachable at a Given Time
 */

// @lc code=start
class Solution {
public:
    bool isReachableAtTime(int sx, int sy, int fx, int fy, int t) {
        if (sx == fx && sy == fy) return t != 1;
        int x = abs(fx-sx);
        int y = abs(fy-sy);
        int minTime = max(x, y);

        return minTime <= t;
    }
};
// @lc code=end
