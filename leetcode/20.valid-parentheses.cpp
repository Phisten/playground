/*
 * @lc app=leetcode id=20 lang=cpp
 *
 * [20] Valid Parentheses
 */

// @lc code=start
class Solution {
public:
    bool isValid(string s) {
        stack<int> stk;
        int n = s.length();

        for (int i = 0; i < n; i++) {
            char cur = s[i];
            if (cur == '(') { stk.push(cur); }
            if (cur == '[') { stk.push(cur); }
            if (cur == '{') { stk.push(cur); }
            else if (stk.size() == 0) return false;
            if (cur == ')') { 
                if (stk.top() == '(') {
                    stk.pop();
                } else {
                    return false;
                }
            }
            if (cur == ']') { 
                if (stk.top() == '[') {
                    stk.pop();
                } else {
                    return false;
                }
            }
            if (cur == '}') { 
                if (stk.top() == '{') {
                    stk.pop();
                } else {
                    return false;
                } 
            }
        }

        return stk.size() == 0;
    }
};
// @lc code=end
