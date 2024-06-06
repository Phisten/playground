/*
 * @lc app=leetcode id=846 lang=cpp
 *
 * [846] Hand of Straights
 */

// @lc code=start
class Solution {
public:
    bool isNStraightHand(vector<int>& hand, int groupSize) {
        sort(hand.begin(), hand.end());

        int n  = hand.size();

        unordered_map<int, int> ht;

        for (int i = 0; i < n; i++) {
            ht[hand[i]]++;
        }

        for (int i = 0; i < n; i++) {
            int start = hand[i];
            if (ht[start] > 0) {
                for (int j = 0 ; j < groupSize; j++) {
                    if (ht[start + j] > 0) {
                        ht[start + j]--;
                    } else return false;
                }
            }
        }

        return true;
    }
};
// @lc code=end

