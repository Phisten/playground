/*
 * @lc app=leetcode id=45 lang=cpp
 *
 * [45] Jump Game II
 */

// @lc code=start
// r = 最大可跳位置
// i = 當前檢查點
// l = 前次最大跳躍

// 每當i > l, ans++(需要跳)
// 每次回圈檢查 r=max(r,nums[i]+i)
class Solution {
public:
    int jump(vector<int>& nums) {
        int n = nums.size();

        int ans = 0;
        int r=0,l=0;
        for (int i = 0; i < n; i++) {
            if (i > l) {
                ans++;
                l = r;
            }
            r = max(r,nums[i]+i);
        }

        return ans;
    }
};
// @lc code=end
