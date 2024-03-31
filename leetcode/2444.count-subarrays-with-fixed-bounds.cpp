/*
 * @lc app=leetcode id=2444 lang=cpp
 *
 * [2444] Count Subarrays With Fixed Bounds
 */

// @lc code=start
/*
1152575 (1,5)
1
11
115 +
>15 +
1152 +
>152 +
11525 +
>1525 +
115257

*/
class Solution
{
public:
    long long countSubarrays(vector<int> &nums, int minK, int maxK)
    {
        int n = nums.size();
        long long ans = 0;
        int minHit = 0, maxHit = 0;
        int lastMaxIdx, lastMinIdx = 0;

        for (int r = 0, l = 0; r < n; r++)
        {
            int v = nums[r];
            if (v > maxK || v < minK)
            {
                l = r + 1;
                minHit = 0;
                maxHit = 0;
                continue;
            }

            if (v == minK)
            {
                minHit++;
                lastMinIdx = r;
            }
            if (v == maxK)
            {
                maxHit++;
                lastMaxIdx = r;
            }
            if (minHit > 0 && maxHit > 0)
                ans += min(lastMaxIdx, lastMinIdx) - l + 1;
        }

        return ans;
    }
};
// @lc code=end
