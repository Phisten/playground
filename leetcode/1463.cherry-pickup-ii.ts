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
    // Base case: If reached out of the last column, return 0 as no more cherries can be picked.
    if (col === cols) return 0;

    // If the result for the current position is already calculated, return the cached value.
    if (dp[col][lastLb][lastRb] >= 0) return dp[col][lastLb][lastRb];

    // Calculate the cherry value at the current path.
    const curValue = lastLb !== lastRb ? grid[col][lastLb] + grid[col][lastRb] : grid[col][lastLb];

    // maximum cherry collection.
    let best = 0;

    // Iterate over possible leftRobot and rightRobot path in the next column.
    for (let l = lastLb - 1; l <= lastLb + 1; l++) {
      for (let r = lastRb - 1; r <= lastRb + 1; r++) {
        // Check if the path are within bounds.
        if (l >= 0 && l < rows && r >= 0 && r < rows) {
          // Recursively calculate the maximum cherry collection for the next column.
          best = Math.max(best, dfs(l, r, col + 1));
        }
      }
    }

    // Cache the result for the current path and return the total cherry collection.
    dp[col][lastLb][lastRb] = curValue + best;
    return dp[col][lastLb][lastRb];
  };

  return dfs(0, rows - 1, 0);
};
// @lc code=end
