/*
 * @lc app=leetcode id=1614 lang=cpp
 *
 * [1614] Maximum Nesting Depth of the Parentheses
 */

// @lc code=start
class Solution
{
public:
    int maxDepth(string s)
    {
        int vps = 0;
        int cur = 0;
        int n = s.length();
        for (int i = 0; i < n; i++)
        {
            if (s[i] == '(')
                cur++;
            if (s[i] == ')' && cur > 0)
                cur--;
            vps = max(vps, cur);
        }
        return vps;
    }
};
// @lc code=end
