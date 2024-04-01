/*
 * @lc app=leetcode id=58 lang=cpp
 *
 * [58] Length of Last Word
 */

// @lc code=start
class Solution {
public:
    int lengthOfLastWord(string s) {
        int n = s.length();

        int r = n - 1;
        for (; r >= 0; r--)
            if (s[r] != ' ')
                break;
        
        int ans = r;
        int l = ans - 1;

        for (; l >= 0; l--)
            if (s[l] == ' ')
                return ans - l;

        return ans + 1;
    }
};
// @lc code=end

