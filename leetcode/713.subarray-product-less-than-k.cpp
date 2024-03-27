/*
 * @lc app=leetcode id=713 lang=cpp
 *
 * [713] Subarray Product Less Than K
 */

// @lc code=start
class Solution
{
public:
    int numSubarrayProductLessThanK(vector<int> &nums, int k)
    {
        int n = nums.size();
        int ans = 0;
        vector<int> cache(nums);
        cache.assign(n, 1);

        int mpl = 1;
        int rangeAns = 0;
        for (int r = 0, l = 0; r < n; r++)
        {
            mpl *= nums[r];
            if (mpl < k)
            {
                ans += r - l + 1;
            }

            while (l <= r && mpl >= k)
            {
                mpl /= nums[l];
                l++;
                if (mpl < k)
                {
                    ans += r - l + 1;
                }
            }
        }

        return ans;
    }
};
// @lc code=end

class Solution_TLE
{
public:
    int numSubarrayProductLessThanK(vector<int> &nums, int k)
    {
        int n = nums.size();
        int ans = 0;
        vector<int> cache(nums);
        cache.assign(n, 1);

        for (int len = 1; len <= n; len++)
        {
            for (int i = 0; i < n - len + 1; i++)
            {
                if (cache[i] < k)
                {
                    cache[i] *= nums[i + len - 1];
                    if (cache[i] < k)
                    {
                        ans++;
                    }
                }
            }
        }
        return ans;
    }
};
