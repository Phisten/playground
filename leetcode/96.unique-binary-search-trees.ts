/*
 * @lc app=leetcode id=96 lang=typescript
 *
 * [96] Unique Binary Search Trees
 */


// 1 2 3  4  5
// 1 2 5 14 42
// @lc code=start
function numTrees(n: number): number {
  if (n <= 2) return n;

  const dp = new Array<number>(n + 1);
  dp[0] = 1;
  dp[1] = 1;
  dp[2] = 2;

  for (let i = 1; i <= n; i++) {
    dp[i] = 0;
    for (let j = 1; j <= i; j++) {
      dp[i] += dp[j - 1] * dp[i - j];
    }
  }

  return dp[n];
};
// @lc code=end
