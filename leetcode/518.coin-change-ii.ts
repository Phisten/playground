/*
 * @lc app=leetcode id=518 lang=typescript
 *
 * [518] Coin Change II
 */
// @lc code=start
export default function change(amount: number, coins: number[]): number {
  const len = coins.length;
  const dp: { [key in number]: number }[] = []

  for (let i = 0; i < len; i++) {
    dp.push({})
    dp[i][0] = 1;
  }

  const dsf = (lastAmount: number, coinIdx: number): number => {
    if (coinIdx < 0 || lastAmount < 0) return 0
    if (dp[coinIdx][lastAmount] !== undefined)
      return dp[coinIdx][lastAmount];

    dp[coinIdx][lastAmount] =
      dsf(lastAmount - coins[coinIdx], coinIdx) +
      dsf(lastAmount, coinIdx - 1);

    return dp[coinIdx][lastAmount];
  }

  const ans = dsf(amount, len - 1);
  return ans;
};

// @lc code=end
