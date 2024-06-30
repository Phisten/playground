class Solution {
public:
    int maximumLength(vector<int>& nums) {
        int k = 2;
        int n = nums.size();
        if (n == 2) return 2;
        
        nums[0] = nums[0] % k;
        int ansDiff = 1;
        int ans0 = nums[0] == 0 ? 1 : 0;
        int ans1 = nums[0] == 1 ? 1 : 0;
        for (int i = 1; i < n; i++) {
            int last = nums[i - 1];
            int cur = nums[i] % k;
            nums[i] = cur;
            
            if (cur != last) {
                ansDiff++;
            }
            if (cur == 0) {
                ans0++;
            }
            if (cur == 1) {
                ans1++;
            }
        }
        
        
        return max(max(ans1, ans0), ansDiff);
    }
};
