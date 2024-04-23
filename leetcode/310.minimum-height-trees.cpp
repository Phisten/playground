/*
 * @lc app=leetcode id=310 lang=cpp
 *
 * [310] Minimum Height Trees
 */

// @lc code=start
class Solution {
public:
    
    vector<int> findMinHeightTrees(int n, vector<vector<int>>& edges) {
        unordered_map<int, vector<int>> ht;
        vector<int> ans;
        int edgeCount = edges.size();

        for (int i = 0; i < edgeCount; i++) {
            ht[edges[i][0]].push_back(edges[i][1]);
            ht[edges[i][1]].push_back(edges[i][0]);
        }

        vector<bool> visited(n);
        function<int(int, int)> testRoot = [&](int root,int deep) -> int {
            if (visited[root] == true) return deep;

            vector<int>& curChild = ht[root];
            visited[root] = true;
            int curMax = 0;
            for (int i = 0; i < curChild.size(); i++) {
                int next = curChild[i];
                if (visited[next] != true)
                    curMax = max(curMax, testRoot(next, deep+1));
            }
            visited[root] = false;

            return curMax;
        };

        int minLevel = n + 1;
        for (int i = 0; i < n; i++) {
            int curLevel = testRoot(i, 1);
            if (curLevel < minLevel) {
                ans.clear();
                minLevel = curLevel;
            }
            if (curLevel <= minLevel) {
                ans.push_back(i);
            }
        }

        return ans;
    }
};
// @lc code=end

