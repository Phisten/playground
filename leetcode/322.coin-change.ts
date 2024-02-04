/*
 * @lc app=leetcode id=322 lang=typescript
 *
 * [322] Coin Change
 */

// @lc code=start
function coinChange(coins: number[], amount: number): number {
  const dp = new Array(amount + 1)

  const dsf = (lastAmount: number) => {
    if (lastAmount === 0)
      return 0
    else if (dp[lastAmount] > 0)
      return dp[lastAmount];

    dp[lastAmount] = amount + 1

    coins.forEach((v: number, i: number) => {
      if (lastAmount - v >= 0) {
        dp[lastAmount] = Math.min(dp[lastAmount], 1 + dsf(lastAmount - v));
      }
    })

    return dp[lastAmount]
  }

  return dsf(amount) > amount ? -1 : dsf(amount)
};
// @lc code=end
