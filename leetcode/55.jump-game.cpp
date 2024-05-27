/*
 * @lc app=leetcode id=55 lang=cpp
 *
 * [55] Jump Game
 */

// @lc code=start
// O(n)
class Solution {
public:
    bool canJump(vector<int>& nums) {
        int n = nums.size();
        if (n == 1) return true;
        int step = 1;
        
        int idx = 0;
        while (step > 0) {
            if (idx == n - 1) return true;
            step--;
            step = max(step, nums[idx]);
            if (step > 0) idx++;
        }

        return false;
    }
};
// @lc code=end

class Solution_DFS {
public:
    bool canJump(vector<int>& nums) {
        int n = nums.size();
        if (n == 1) return true;
        
        bool ans = false;
        function<void(int)> dfs = [&](int index) {
            if (index >= n - 1) ans = true;
            if (ans) return;

            int maxStep = nums[index];

            for (int i = 1; i <= maxStep; i++) {
                dfs(maxStep-i+1+ index);
            }

            // deduplication
            nums[index] = 0;
        };
        dfs(0);

        return ans;
    }
};
