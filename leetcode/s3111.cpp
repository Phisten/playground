class Solution {
public:
    int minRectanglesToCoverPoints(vector<vector<int>>& points, int w) {
        // xy inc order
        sort(points.begin(), points.end(), [](const vector<int>& a, const vector<int>& b) {
            return a[0] < b[0] || (a[0] == b[0] && (a[1] < b[1]));
        });

        int n = points.size();

        // ht: x to first index + 1
        unordered_map<int, int> ht;
        // ht keys
        vector<int> px;
        for (int i = 0; i < n; i++) {
            if (ht[points[i][0]] == 0) {
                ht[points[i][0]] = i + 1;
                px.push_back(points[i][0]);
                cout << points[i][0] << endl;
            }
        }

        int ans = 0;
        int lastX = -w-1;
        for (int i = 0; i < px.size(); i++) { 
            if (px[i] - lastX > w) {
                ans++;
                lastX = px[i];
            }
        }

        return ans;
    }
};
