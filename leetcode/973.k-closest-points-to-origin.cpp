/*
 * @lc app=leetcode id=973 lang=cpp
 *
 * [973] K Closest Points to Origin
 */

// @lc code=start
class Solution {
public:
    struct PointWithDist {
        vector<int> point;
        int dis;
    };

    struct Compare {
        bool operator()(const PointWithDist& a, const PointWithDist& b) {
            return a.dis > b.dis;
        }
    };

    vector<vector<int>> kClosest(vector<vector<int>>& points, int k) {
        int n = points.size();
        priority_queue<PointWithDist, vector<PointWithDist>, Compare> pq;

        for (int i = 0; i < n; i++) {
            PointWithDist p;
            p.point = points[i];
            p.dis = points[i][0]*points[i][0] + points[i][1]*points[i][1];

            pq.push(p);
        }


        vector<vector<int>> ans;
        for (int i = 0; i < k; i++) {
            auto& a = pq.top();
            ans.push_back(a.point);
            pq.pop();
        }
        return ans;
    }
};
// @lc code=end

