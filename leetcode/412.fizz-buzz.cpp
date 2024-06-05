/*
 * @lc app=leetcode id=412 lang=cpp
 *
 * [412] Fizz Buzz
 */

// @lc code=start
class Solution {
public:
    vector<string> fizzBuzz(int n) {
        vector<string> ans;

        for (int i = 1; i <= n; i++) {
            ostringstream s;
            if (i % 3 == 0) s << "Fizz";
            if (i % 5 == 0) s << "Buzz";
            if (i % 5 != 0 && i % 3 != 0) s << to_string(i);

            ans.push_back(s.str());
        }

        return ans;
    }
};
// @lc code=end

