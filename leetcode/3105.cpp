class Solution
{
public:
  int longestMonotonicSubarray(vector<int> &nums)
  {
    int n = nums.size();
    int ans = 0;

    unordered_map<int, int> hti;
    unordered_map<int, int> htd;

    int lastR = 0;
    for (int r = 0, li = 0, ld = 0; r < n; r++)
    {
      int valR = nums[r];

      if (valR > lastR)
      {
        ans = max(ans, r - li + 1);
      }
      else
      {
        li = r;
      }
      if (valR < lastR)
      {
        ans = max(ans, r - ld + 1);
      }
      else
      {
        ld = r;
      }
      lastR = valR;
    }
    return ans;
  }
};
