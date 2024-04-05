/*
 * @lc app=leetcode id=1759 lang=cpp
 *
 * [1759] Count Number of Homogenous Substrings
 */

// @lc code=start
class Solution
{
public:
    int countHomogenous(string s)
    {
        int n = s.length();
        int ans = 0;

        for (int r = 0, l = 0; r < n; r++)
        {
            if (s[r] == s[l])
            {
                ans = (ans + r - l + 1) % 1000000007;
            }
            else
            {
                l = r;
                ans = (ans + 1) % 1000000007;
            }
        }

        return ans;
    }
};
// @lc code=end
