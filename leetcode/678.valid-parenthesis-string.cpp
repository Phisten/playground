/*
 * @lc app=leetcode id=678 lang=cpp
 *
 * [678] Valid Parenthesis String
 */

// @lc code=start
class Solution
{
public:
    bool checkValidString(string s)
    {
        int n = s.length();
        bool stable = false;
        while (stable == false)
        {
            stable = true;
            for (int l = 0, r = 0; l < n; l++)
            {
                if (s[l] == '(')
                {
                    r = l + 1;
                    while (r < n && s[r] != '(')
                    {
                        if (s[r] == ')')
                        {
                            s[l] = '.';
                            s[r] = '.';
                            stable = false;
                            break;
                        }

                        r++;
                    }
                }
            }
        }

        int any = 0, l = 0, spareL = 0;

        for (int i = 0; i < n; i++)
        {
            if (s[i] == '(')
            {
                l++;
            }
            if (s[i] == ')')
            {
                if (any > 0)
                    any--;
                else
                    return false;
            }
            if (s[i] == '*')
            {
                if (l > 0)
                    l--;
                else
                    any++;
            }
        }
        return l == 0;
    }
};
// @lc code=end
