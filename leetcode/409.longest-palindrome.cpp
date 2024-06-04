/*
 * @lc app=leetcode id=409 lang=cpp
 *
 * [409] Longest Palindrome
 */

// @lc code=start
class Solution {
public:
    int longestPalindrome(string s) {
        int n = s.length();
        int hasOdd = 0;
        unordered_map<char, int> ht;

        for (int i = 0; i < n; i++) {
            ht[s[i]]++;
        }

        int ans = 0;
        for (auto& [k, v] : ht) {
            int mod = v % 2;
            if (mod == 1) hasOdd = 1;
            ans += v - mod;
        }

        return ans + hasOdd;
    }
};
// @lc code=end

