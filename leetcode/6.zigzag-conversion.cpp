/*
 * @lc app=leetcode id=6 lang=cpp
 *
 * [6] Zigzag Conversion
 */

// @lc code=start
class Solution
{
public:
    string convert(string s, int numRows)
    {
        vector<string> mat(numRows);
        int n = s.length();

        for (int i = 0, row = 0, rowAdd = 1; i < n; i++)
        {
            mat[row] += s[i];

            row = min(row + rowAdd, numRows - 1);
            if (row == 0)
                rowAdd = 1;
            else if (row == numRows - 1)
                rowAdd = -1;
        }

        for (int i = 1; i < numRows; i++)
            mat[0].append(mat[i]);

        return mat[0];
    }
};
// @lc code=end
