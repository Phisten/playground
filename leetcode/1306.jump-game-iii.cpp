/*
 * @lc app=leetcode id=1306 lang=cpp
 *
 * [1306] Jump Game III
 */

// @lc code=start
class Solution {
public:
    bool canReach(vector<int>& arr, int start) {
        int n = arr.size();
        if (n == 1) return true;

        deque<int> dq;
        dq.push_back(start);

        while (dq.size() > 0) {
            int cur = dq.front();
            dq.pop_front();

            if (cur >= n || cur < 0)
                continue;
            else if (arr[cur] == 0) {
                return true;
            } 

            int curValue = arr[cur];
            if (curValue != -1) {
                arr[cur] = -1;
                dq.push_back(cur + curValue);
                dq.push_back(cur - curValue);
            }
        }

        return false;
    }
};
// @lc code=end


class Solution_dfs {
public:
    bool canReach(vector<int>& arr, int start) {
        int n = arr.size();
        if (n == 1) return true;

        bool ans = false;
        function<void(int)> dfs = [&](int cur) {
            if (cur >= n || cur < 0)
                return;
            else if (arr[cur] == 0 || ans) {
                ans = true;
                return;
            } 

            int curValue = arr[cur];
            if (curValue != -1) {
                arr[cur] = -1;
                dfs(cur + curValue);
                dfs(cur - curValue);
            }
        };
        dfs(start);

        return ans;
    }
};
