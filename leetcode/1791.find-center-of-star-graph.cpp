/*
 * @lc app=leetcode id=1791 lang=cpp
 *
 * [1791] Find Center of Star Graph

# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
The problem is to find the center of a star graph. A star graph is a graph where one central node is connected to all other nodes, and those nodes are not connected to each other.

The central node of the star graph will always appear in every edge of the given edge list because every other node is only connected to the center.

# Approach
<!-- Describe your approach to solving the problem. -->
1. Compare the first node of the first edge with both nodes of the second edge.

2. If the first node of the first edge matches either node in the second edge, it is the central node.

3. Otherwise, the second node of the first edge must be the central node.


 */

// @lc code=start
class Solution {
public:
    // Function to find the center of the star graph
    int findCenter(vector<vector<int>>& edges) {
        // The center node will appear in the first two edges. 
        // Check if the first element of the first edge is the same as any element in the second edge.
        // If it is, then the first element of the first edge is the center.
        // Otherwise, the second element of the first edge is the center.
        return edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] ? edges[0][0] : edges[0][1];
    }
};
// @lc code=end


class Solution {
public:
    int findCenter_first(vector<vector<int>>& edges) {
        unordered_map<int, int> ht;
        int n = edges.size();

        for (int i = 0; i < n; i++) {
            ht[edges[i][1]]++;
            ht[edges[i][0]]++;

            if (ht[edges[i][0]] > 1) return edges[i][0];
            if (ht[edges[i][1]] > 1) return edges[i][1];
        }

        return 1;
    }
};
