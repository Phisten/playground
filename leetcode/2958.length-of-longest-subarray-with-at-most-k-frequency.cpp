/*
 * @lc app=leetcode id=2958 lang=cpp
 *
 * [2958] Length of Longest Subarray With at Most K Frequency
 */

// @lc code=start
class Solution {
public:
    int maxSubarrayLength(vector<int>& nums, int k) {
        int n = nums.size();
        unordered_map<int, int> ht;
        int valid = 0;
        int subLen = 0;
        int ans = 0;

        for (int r = 0, l = 0; r < n; r++) {
            int v = nums[r];
            ht[v] += 1;

            if (ht[v] == 1) {
                subLen++;
                valid++;
            } else if (ht[v] == k+1) {
                valid--;
            }
            if (subLen == valid) {
                ans = max(ans, r - l + 1);
            }

            while (l <= r && subLen != valid) {
                int vL = nums[l];
                ht[vL] -= 1;
                l++;
                
                if (ht[vL] == 0) {
                    subLen--;
                    valid--;
                } else if (ht[vL] == k) {
                    valid++;
                }
                if (subLen == valid) {
                    ans = max(ans, r - l + 1);
                }
            }
        }
        return ans;
    }
};
// @lc code=end
