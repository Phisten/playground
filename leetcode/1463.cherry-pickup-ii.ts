/*
 * @lc app=leetcode id=1463 lang=typescript
 *
 * [1463] Cherry Pickup II
 */

// @lc code=start
function cherryPickup(grid: number[][]): number {
  const cols = grid.length, rows = grid[0].length;

  // store calculated values for [col Index][left Robot Selected Path][right Robot Selected Path]
  const dp: number[][][] = grid.map(() =>
    Array.from({ length: rows }, () => Array.from<number>({ length: rows }).fill(-1)));

  const dfs = (lastLb, lastRb, col): number => {


    if (col === cols) return 0;
    if (dp[col][lastLb][lastRb] >= 0) return dp[col][lastLb][lastRb];

    const curValue = lastLb !== lastRb ? grid[col][lastLb] + grid[col][lastRb] : grid[col][lastLb];
    let best = 0
    for (let l = lastLb - 1; l <= lastLb + 1; l++) {
      for (let r = lastRb - 1; r <= lastRb + 1; r++) {
        if (l >= 0 && l < rows && r >= 0 && r < rows) {
          best = Math.max(best, dfs(l, r, col + 1));
        }
      }
    }

    // Cache the result for the current position and return the total cherry collection.
    dp[col][lastLb][lastRb] = curValue + best;
    return curValue + best;
  };

  const ans = dfs(0, rows - 1, 0);
  return ans;
};
// @lc code=end
