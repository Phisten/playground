
describe('863. All Nodes Distance K in Binary Tree', () => {

})

class TreeNode {
  val: number
  left: TreeNode | null
  right: TreeNode | null
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
  }
}

function distanceK(root: TreeNode | null, target: TreeNode | null, k: number): number[] {
  const parent: { [key in number]: TreeNode } = {}

  // find parent 
  const dfs = (node, p) => {
    if (!node) return;

    parent[node.val] = p;

    dfs(node.left, node);
    dfs(node.right, node);
  };
  dfs(root, null);

  const ans = [];
  const visited = {}
  const recDist = (node, lastK) => {
    if (!node || visited[node.val]) return;
    if (lastK === 0)
      ans.push(node.val);

    visited[node.val] = true;

    recDist(node.left, lastK - 1);
    recDist(node.right, lastK - 1);
    recDist(parent[node.val], lastK - 1);
  }
  recDist(target, k)

  return ans;
};
