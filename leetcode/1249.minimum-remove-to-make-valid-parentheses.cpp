/*
 * @lc app=leetcode id=1249 lang=cpp
 *
 * [1249] Minimum Remove to Make Valid Parentheses
 */

// @lc code=start
class Solution {
public:
    string minRemoveToMakeValid(string s) {
        stringstream ans;

        int pair = 0;
        int pairTemp = 0;
        int n = s.length();


        for (int i = 0; i < n; i++) {
            if (s[i] == '(')
                pairTemp++;
            if (s[i] == ')')
                if (pairTemp > 0) {
                    pairTemp--;
                    pair++;
                } else
                    s[i] = '.';
        }

        int l = 0,r = 0;
        for (int i = 0; i < n; i++) {
            if (s[i] == '(') {
                l++;
                if (l > pair) continue;
            }
            if (s[i] == ')') {
                r++;
                if (r > pair) continue;
            }
            if (s[i] != '.')
                ans << s[i];
        }

        return ans.str();
    }
};
// @lc code=end
