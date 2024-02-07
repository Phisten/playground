/*
 * @lc app=leetcode id=451 lang=typescript
 *
 * [451] Sort Characters By Frequency
 */

// @lc code=start
function frequencySort(s: string): string {
  const ht: { [key in string]: number } = {}

  for (let item of s) {
    if (!ht[item]) ht[item] = 0;
    ht[item] += 1
  }

  const arr: { char: string, count: number }[] = Object.keys(ht).map((v, i) => ({ char: v, count: ht[v] }))
  const ordered = arr.sort((a, b) => b.count - a.count)

  let ans = ordered.map(element =>
    element.char.repeat(element.count)
  ).join('');

  return ans;
};
// @lc code=end
