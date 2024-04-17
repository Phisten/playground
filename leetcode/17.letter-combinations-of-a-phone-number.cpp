/*
 * @lc app=leetcode id=17 lang=cpp
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
class Solution
{
public:
    vector<string> letterCombinations(string digits)
    {
        unordered_map<char, string> ht;

        ht['2'] = "abc";
        ht['3'] = "def";
        ht['4'] = "ghi";
        ht['5'] = "jkl";
        ht['6'] = "mno";
        ht['7'] = "pqrs";
        ht['8'] = "tuv";
        ht['9'] = "wxyz";

        int n = digits.length();
        vector<string> ans;

        function<void(string, int)> dfs = [&](string lastStr, int idx)
        {
            if (idx == n)
            {
                if (lastStr != "")
                    ans.push_back(lastStr);
                return;
            }

            int nextIdx = idx + 1;
            string charList = ht[digits[idx]];
            for (int i = 0; i < charList.length(); i++)
            {
                dfs(lastStr + charList[i], nextIdx);
            }
        };
        dfs("", 0);

        return ans;
    }
};
// @lc code=end
