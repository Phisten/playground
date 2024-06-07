/*
 * @lc app=leetcode id=14 lang=cpp
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        int n = strs.size();
        if (n == 1) return strs[0];

        int j = 0;
        int nn = strs[0].length();

        while (j < nn) {
            char c = strs[0][j];

            for (auto& str : strs)
                if (j >= str.length() || c != str[j])
                    return strs[0].substr(0, j);
                
            j++;
        }

        return strs[0];
    }
};
// @lc code=end

