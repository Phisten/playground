class Solution {
public:
    vector<vector<int>> subsets(vector<int>& nums) {
        vector<vector<int>> ans;

        int n = nums.size();
        int bitN = pow(2, n);

        for (int i = 0; i < bitN; i++) {
            vector<int> curAns;
            for (int j = 0; j < n; j++) {
                cout << (i >> j) << endl;
                if (((i >> j) & 1) == 1) {
                    curAns.push_back(nums[j]);
                }
            }
            ans.push_back(curAns);
        }

        return ans;
    }
};