/*
 * @lc app=leetcode id=279 lang=typescript
 *
 * [279] Perfect Squares
 */

// @lc code=start
function numSquares(n: number): number {
  if (n <= 0) return 0;
  if (n === 1) return 1;

  const squares: number[] = Array.from(
    { length: Math.floor(Math.sqrt(n)) + 1 },
    (_, i) => i * i,
  );

  const dp: number[] = Array.from({ length: n });

  const dfs = (lastNum): number => {
    if (lastNum < 4) return lastNum;
    if (dp[lastNum]) return dp[lastNum];

    let minNextRes = lastNum - 1;
    for (let i = 1; i < squares.length; i++) {
      const item = squares[squares.length - i];
      if (lastNum > item)
        minNextRes = Math.min(dfs(lastNum - item), minNextRes)
      else if (lastNum === item) {
        minNextRes = 0;
        break;
      }
    }

    dp[lastNum] = 1 + minNextRes;

    return 1 + minNextRes;
  }

  return dfs(n)
};
// @lc code=end
