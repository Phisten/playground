/*
 * @lc app=leetcode id=386 lang=typescript
 *
 * [386] Lexicographical Numbers
 */

// @lc code=start
function largestDivisibleSubset(nums: number[]): number[] {
  const length = nums.length

  nums.sort((a, b) => a - b);

  const dp: number[] = Array.from<number>({ length }).fill(1);
  dp[0] = 1;

  const linked: number[] = Array.from<number>({ length });
  let maxSubset = 1;
  let headIdx = 0;

  for (let i = 1; i < length; i++) {
    for (let j = 0; j < i; j++) {
      if (nums[i] % nums[j] == 0) {
        if (dp[j] + 1 > dp[i]) {
          dp[i] = dp[j] + 1;
          linked[i] = j;

          if (dp[i] > maxSubset) {
            headIdx = i;
            maxSubset = dp[i];
          }
        }
      }
    }
  }

  const ans = [];
  while (nums[headIdx]) {
    ans.push(nums[headIdx])
    headIdx = linked[headIdx];
  }

  return ans;
};
// @lc code=end
