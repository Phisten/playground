/*
 * @lc app=leetcode id=11 lang=cpp
 *
 * [11] Container With Most Water
 */

// @lc code=start
class Solution
{
public:
    int maxArea(vector<int> &height)
    {
        int n = height.size();
        int l = 0, r = n - 1;
        int ans = 0;

        while (l < r)
        {
            ans = max(ans, (r - l) * min(height[l], height[r]));

            if (height[l] < height[r])
            {
                l++;
            }
            else
            {
                r--;
            }
        }

        return ans;
    }
};
// @lc code=end

class Solution_bruteForce_TLE
{
public:
    int maxArea(vector<int> &height)
    {
        int n = height.size();

        int ans = 0;
        for (int i = 0; i <= n; i++)
        {
            for (int j = i + 1; j < n; j++)
            {
                ans = max(ans, (j - i) * min(height[i], height[j]));
            }
        }

        return ans;
    }
};
