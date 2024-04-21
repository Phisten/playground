/*
 * @lc app=leetcode id=1971 lang=cpp
 *
 * [1971] Find if Path Exists in Graph
 */

// @lc code=start
class Solution {
public:
    bool validPath(int n, vector<vector<int>>& edges, int source, int destination) {
        unordered_map<int, vector<int>> ht;

        for (int i = 0; i < edges.size(); i++) {
            ht[edges[i][0]].push_back(edges[i][1]);
            ht[edges[i][1]].push_back(edges[i][0]);
        }

        vector<int> used(n);
        bool ans = false;
        function<void(int)> dfs = [&](int i) {
            if (used[i] == 1 || ans) return;
            
            if (i == destination) {
                ans = true;
                return;
            }

            used[i] = 1;
            
            auto& curHt = ht[i];
            for (int j = 0; j < curHt.size(); j++) {
                if (used[curHt[j]] == 0) dfs(curHt[j]);
            }
        };

        dfs(source);

        return ans;
    }
};
// @lc code=end

