/*
 * @lc app=leetcode id=1544 lang=cpp
 *
 * [1544] Make The String Great
 */

// @lc code=start
class Solution {
public:
    string makeGood(string s) {
        int shift = 'a' - 'A';

        bool run = true;
        while (run) {
            int n = s.length();
            run = false;
            stringstream tmp;

            for (int i = 0; i < n; i++) {
                if (s[i] == ' ') continue;
    
                if (i+1 < n && (s[i] == s[i+1] + shift || s[i] == s[i+1] - shift)) {
                    s[i] = ' ';
                    s[i+1] = ' ';
                    run = true;
                } else {
                    tmp << s[i];
                }
            }
            s = tmp.str();
        }
        
        return s;
    }
};
// @lc code=end

