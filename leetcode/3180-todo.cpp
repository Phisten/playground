class Solution {
public:
    int maxTotalReward(vector<int>& rewardValues) {
      int n = rewardValues.size();
      sort(rewardValues.begin(),rewardValues.end());

      function<int(int, int)> dfs = [&](int limit, int lastIdx) {
        int lastVal = rewardValues[lastIdx];

        int maxRes = 0;
        for (int i = lastIdx - 1; i >= 0 ; i--) {
          int cur = rewardValues[i];
          if (cur < lastVal) { // 跳過重複
            if (cur < limit) {
              int res = dfs(min(lastVal - cur, limit - cur), i);
              maxRes = max(maxRes, res);
            }
          }
        }
        return maxRes + rewardValues[lastIdx];
      };

      int ans = dfs(2147483647 ,n-1);
      
      return ans;
    }
};


/*
[1,1,3,3] = 4

[1,6,4,3,2] = 11
[1,2,3,4,6]= 11
6+4

[1,100,101,200,302] = 
  1+100+200=301 +302=603
  1+101+200=302

[1,100,101,200,201,302] = 
  1+100+200=301 +302=603
  1+101+200=302

302+201+200=703

6+


[2,100,101,200] = 
2+100=102 +200=302
2+101=103 +200=303
2+200=202

100+101=201
101+200=301
200


[2,8,9,10] = 
19
2+8+=10
*/
