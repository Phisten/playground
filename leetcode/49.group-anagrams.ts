/*
 * @lc app=leetcode id=49 lang=typescript
 *
 * [49] Group Anagrams
 */

// @lc code=start
function groupAnagrams(strs: string[]): string[][] {
  const ht: { [key in string]: string[] } = {}

  strs.forEach(element => {
    const sorted = element.split('').sort().join();
    if (!ht[sorted]) ht[sorted] = []

    ht[sorted].push(element)
  });

  const ans = Object.values(ht);
  return ans;
};
// @lc code=end
