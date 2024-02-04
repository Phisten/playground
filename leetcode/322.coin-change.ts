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
    else if (lastAmount < 0) {
      return -1
    }

    if (dp[lastAmount] > 0) {
      return dp[lastAmount];
    }

    let res = amount + 1

    coins.forEach((v: number, i: number) => {
      const cur = dsf(lastAmount - v)
      if (cur !== -1) {
        res = Math.min(res, 1 + cur);
      }
    })

    dp[lastAmount] = res
    return res
  }

  const count = dsf(amount)
  return count > amount ? -1 : count
};
// @lc code=end
