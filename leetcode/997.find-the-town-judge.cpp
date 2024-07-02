/*
 * @lc app=leetcode id=997 lang=cpp
 *
 * [997] Find the Town Judge
 */

// @lc code=start
class Solution {
public:
    int findJudge(int n, vector<vector<int>>& trust) {
        unordered_map<int, int> ht;  // Hash table to track trust counts

        int m = trust.size();  // Number of trust relationships

        // Traverse the trust relationships
        for (int i = 0; i < m; i++) {
            int cur = trust[i][1];  // The person who is trusted
            ht[cur]++;  // Increment the trust count for the person who is trusted
            ht[trust[i][0]] = -n;  // Mark the person who trusts someone as not the judge
        }

        // Check each person from 1 to n to find the judge
        for (int i = 1; i <= n; i++) {
            if (ht[i] == n - 1) return i;  // The judge should be trusted by exactly n-1 people and trust no one
        }

        return -1;  // If no judge is found, return -1
    }
};
// @lc code=end
