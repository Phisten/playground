/*
 * @lc app=leetcode id=331 lang=typescript
 *
 * [331] Verify Preorder Serialization of a Binary Tree
 */

// @lc code=start
function isValidSerialization(preorder: string): boolean {
  const values = preorder.split(',');
  const length = values.length;

  if (length === 1) return values[0] === '#';

  let rightStack = 0;
  let nodeType: (-1 | 1 | 0) = 0;
  for (let i = 0; i < length; i++) {
    if (rightStack < 0) return false;

    switch (nodeType) {
      case -1: // left
        rightStack++;
        if (values[i] === '#')
          nodeType = 1;
        break;
      case 1: // right
        rightStack--;
        if (values[i] !== '#')
          nodeType = -1;
        break;
      default: // root
        if (values[i] === '#')
          return false;
        else
          nodeType = -1;
    }
  }

  return rightStack === 0;
};
// @lc code=end
