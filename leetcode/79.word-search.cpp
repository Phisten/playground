/*
 * @lc app=leetcode id=79 lang=cpp
 *
 * [79] Word Search
 */

// @lc code=startclass Solution {
public:
bool exist(vector<vector<char>> &board, string word)
{
    int m = board.size();
    int n = board[0].size();
    int len = word.length();

    function<bool(int, int, int)> match = [&](int i, int j, int idx)
    {
        if (idx == len)
            return true;
        if (i < 0 || i >= m || j < 0 || j >= n)
            return false;
        bool res = false;

        if (board[i][j] == word[idx])
        {
            board[i][j] = 0;
            res = match(i - 1, j, idx + 1) ||
                  match(i, j - 1, idx + 1) ||
                  match(i + 1, j, idx + 1) ||
                  match(i, j + 1, idx + 1);
            board[i][j] += word[idx];
        }
        return res;
    };

    for (int i = 0; i < m; i++)
        for (int j = 0; j < n; j++)
            if (match(i, j, 0))
                return true;

    return false;
}
}
;
// @lc code=end
