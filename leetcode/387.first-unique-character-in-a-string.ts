/*
 * @lc app=leetcode id=387 lang=typescript
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
function firstUniqChar(s: string): number {
  let ht = {}
  for (let item of s)
    ht[item] = ht[item] ? ht[item] + 1 : 1

  for (let i = 0; i < s.length; i++)
    if (ht[s[i]] === 1)
      return i;

  return -1;
};
// @lc code=end
