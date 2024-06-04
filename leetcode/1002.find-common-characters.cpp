/*
 * @lc app=leetcode id=1002 lang=cpp
 *
 * [1002] Find Common Characters
 */

// @lc code=start
class Solution {
public:
    vector<string> commonChars(vector<string>& words) {
        int n = words.size();
        vector<string> ans;
        vector<unordered_map<char, int>> ht(n);

        for (int i = 0; i < n ; i++) {
            int nn = words[i].size();

            for (int j = 0; j < nn; j++) {
                ht[i][words[i][j]]++;
            }
        }

        int nn = words[0].size();
        for (int i = 0; i < nn; i++) {
            char c = words[0][i];
            if (ht[0][c] > 0) {
                int cCount = ht[0][c];
                for (int j = 1; j < n; j++)
                    cCount = min(cCount, ht[j][c]);

                for (int j = 0; j < cCount; j++) {
                    std::string str(1, c);
                    ans.push_back(str);
                }

                ht[0][c] = 0;
            }
        }

        return ans;
    }
};
// @lc code=end

