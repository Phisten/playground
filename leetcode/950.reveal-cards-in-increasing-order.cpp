/*
 * @lc app=leetcode id=950 lang=cpp
 *
 * [950] Reveal Cards In Increasing Order
 */

// @lc code=start
/*
2 3 5 7 B D G  increasing order
2 D 3 B 5 G 7  ans

1 2 3 4 5 6 7
1 6 2 5 3 7 4 

1.照順序每填一個數字後就skip一個位置
2.填完一輪再回到最前面空位重複1
*/

class Solution {
public:
    vector<int> deckRevealedIncreasing(vector<int>& deck) {
        int n = deck.size();
        vector<int> ans(n, 0);
        sort(deck.begin(), deck.end());

        for (int top = 0, skip = 0; top < n; ) {
            for (int j = 0; j < n; j++) {
                if (0 == ans[j]) {
                    if (0 == skip) {
                        ans[j] = deck[top];
                        top++;
                        skip = 1;
                    } else
                        skip = 0;
                }
            }
        }

        return ans;
    }
};
// @lc code=end
