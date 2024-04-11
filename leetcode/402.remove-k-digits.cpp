/*
 * @lc app=leetcode id=402 lang=cpp
 *
 * [402] Remove K Digits
 */

// @lc code=start
class Solution
{
public:
    string removeKdigits(string num, int k)
    {
        int n = num.length();
        if (k == n)
            return "0";

        string ans;
        ans.push_back(num[0]);

        int need = n - k - 1;
        // cout << "ans: " << ans << endl;

        for (int i = 1; i < n; i++)
        {
            int v = num[i];

            while (ans.length() > 0 && ans.back() > v && k > 0)
            {
                ans.pop_back();
                need++;
                k--;
            }

            if (ans.length() > 0 || v != '0')
            {
                ans.push_back(v);
            }
            need--;

            // cout << "ans: " << ans << endl;
        }

        if (need < 0)
            for (int i = need; i < 0; i++)
                if (ans.length() > 0)
                    ans.pop_back();
                else
                    break;

        return ans == "" ? "0" : ans;
    }
};
// @lc code=end
