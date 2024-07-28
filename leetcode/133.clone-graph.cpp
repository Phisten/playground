/*
 * @lc app=leetcode id=133 lang=cpp
 *
 * [133] Clone Graph
 */

// @lc code=start
/*
/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> neighbors;
    Node() {
        val = 0;
        neighbors = vector<Node*>();
    }
    Node(int _val) {
        val = _val;
        neighbors = vector<Node*>();
    }
    Node(int _val, vector<Node*> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
};
*/

class Solution
{
    public:
        Node* cloneGraph(Node *node)
        {
            unordered_map<int, Node*> ht;
            Node *root = new Node();
            Node *last = root;

            function<void(Node*)> dfs =[&](Node *cur) {
                if (cur == nullptr) return;
                if (ht.count(cur->val) > 0) return;

                ht[cur->val] = new Node(cur->val);

                int n = cur->neighbors.size();
                for (int i = 0; i < n; i++) {
                    dfs(cur->neighbors[i]);
                    ht[cur->val]->neighbors.push_back(ht[cur->neighbors[i]->val]);
                }
            };
            dfs(node);

            return nullptr == node ? nullptr : ht[node->val];
        }
};
// @lc code=end
