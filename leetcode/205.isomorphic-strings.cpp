/*
 * @lc app=leetcode id=205 lang=cpp
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
class Solution {
public:
    bool isIsomorphic(string s, string t) {
        unordered_map<char, char> ht_ts;
        unordered_map<char, char> ht_st;

        int n = s.length();

        for (int i = 0; i < n; i++) {
            if (ht_ts.count(t[i]) == 0) ht_ts[t[i]] = s[i];
            if (ht_ts[t[i]] != s[i]) return false;

            if (ht_st.count(s[i]) == 0) ht_st[s[i]] = t[i];
            if (ht_st[s[i]] != t[i]) return false;
        }

        return true;
    }
};
// @lc code=end

