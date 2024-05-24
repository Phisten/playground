/*
 * @lc app=leetcode id=131 lang=cpp
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
class Solution {
public:
    vector<vector<string>> partition(string s) {
        vector<vector<string>> ans;
        int n = s.length();

        vector<string> cache;
        // 分割數
        function<void(string,int)> dfs = [&](string s, int start) {
            if (start >= n) return;
            for (int l = 1; l <= n-start; l++) {

                int isPalind = true;
                for (int i = start, j = start + l - 1; i <= start + l / 2; i++, j--) {
                    if (s[i] != s[j]) isPalind = false;
                }

                if (isPalind) {
                    cache.push_back(s.substr(start, l));

                    if (start + l == n) {
                        ans.push_back(cache);
                    }

                    dfs(s, start + l);
                    cache.pop_back();
                }
            }
        };
        dfs(s, 0);

        return ans;
    }
};
// @lc code=end

