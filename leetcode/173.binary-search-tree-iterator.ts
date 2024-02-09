/*
 * @lc app=leetcode id=173 lang=typescript
 *
 * [173] Binary Search Tree Iterator
 */

// @lc code=start
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

class BSTIterator {
    nodes: TreeNode[] = [];
    idx: number = 0;

    constructor(root: TreeNode) {
        this.nodes = [];
        this.idx = 0;

        let dsf = (node: TreeNode | null) => {
            if (!node) return;

            if (this.nodes.length > 0 && this.nodes[this.idx].val > node.val) {
                this.idx = this.nodes.length;
            }

            this.nodes.push(node)
            dsf(node.left);
            dsf(node.right);
        }
        dsf(root);

        this.idx--;

        console.log('this.nodes', this.nodes);
    }

    next(): number {
        this.idx += 1;
        const val = this.nodes[this.idx]?.val;
        return val;
    }

    hasNext(): boolean {
        return this.nodes.length > this.idx + 1;
    }
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * var obj = new BSTIterator(root)
 * var param_1 = obj.next()
 * var param_2 = obj.hasNext()
 */
// @lc code=end
