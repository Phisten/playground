/*
 * @lc app=leetcode id=992 lang=cpp
 *
 * [992] Subarrays with K Different Integers
 */

// @lc code=start
class Solution
{
public:
    int subarraysWithKDistinct(vector<int> &nums, int k)
    {
        int n = nums.size();
        unordered_map<int, int> ht;
        int vaild = 0;
        int eqBgK = 0;

        for (int i = 0, l = 0; i < n; i++)
        {
            int v = nums[i];
            if (ht[v] == 0)
            {
                vaild++;
            }
            if (vaild >= k)
            {
                eqBgK += n - i;
            }
            ht[v]++;

            while (l <= i && vaild >= k)
            {
                int v = nums[l];
                ht[v]--;
                if (ht[v] == 0)
                {
                    vaild--;
                }
                if (vaild >= k)
                {
                    eqBgK += n - i;
                }

                l++;
            }
        }

        ht.clear();
        vaild = 0;
        int bgK = 0;

        for (int i = 0, l = 0; i < n; i++)
        {
            int v = nums[i];
            if (ht[v] == 0)
            {
                vaild++;
            }
            if (vaild > k)
            {
                bgK += n - i;
            }
            ht[v]++;

            while (l <= i && vaild > k)
            {
                int v = nums[l];
                ht[v]--;
                if (ht[v] == 0)
                {
                    vaild--;
                }
                if (vaild > k)
                {
                    bgK += n - i;
                }

                l++;
            }
        }

        return eqBgK - bgK;
    }
};
// @lc code=end
