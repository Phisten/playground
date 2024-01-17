/*
 * @lc app=leetcode id=1207 lang=cpp
 *
 * [1207] Unique Number of Occurrences
 */

// @lc code=start
#include <vector>
#include <unordered_map>
using namespace std;

class Solution
{
public:
    bool uniqueOccurrences(vector<int> &arr)
    {
        unordered_map<int, int> ht;
        unordered_map<int, int> htCount;

        for (auto &&i : arr)
            ht[i]++;

        for (auto &&pair : ht)
            if (htCount[pair.second] != 1)
                htCount[pair.second]++;
            else
                return false;

        return true;
    }
};
// @lc code=end
