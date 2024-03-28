/*
 * @lc app=leetcode id=5 lang=cpp
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
class Solution
{
public:
    string longestPalindrome(string s)
    {
        int n = s.length();
        int ansI = 0, ansLen = 1;
        int ansL = 0, ansR = 0;

        function<void(int, int)> isPalindromic = [&](int l, int r)
        {
            int len = r - l + 1;

            if (l >= 0 && r < n && s[l] == s[r])
            {
                if (ansLen < len)
                {
                    ansLen = len;
                    ansL = l;
                    ansR = r;
                }

                isPalindromic(l - 1, r + 1);
            }
        };

        for (int i = 0; i < n; i++)
        {
            // odd
            isPalindromic(i, i);

            // even
            isPalindromic(i, i + 1);
        }

        return s.substr(ansL, ansR - ansL + 1);
    }
};
// @lc code=end
