/*
 * @lc app=leetcode id=402 lang=cpp
 *
 * [402] Remove K Digits
 */

// @lc code=start
class Solution {
public:
    string removeKdigits(string num, int k) {
        int n = num.length();
        stringstream ss;
        cout << n << endl;
        // cout << "k=" << k << "," << num << endl;

        for (int i = 0; i < n && k > 0;) {
            if (num[i] >= '0') {
                int maxNumIdx = i;
                int lastNumIdx = i;
                for (int j = i + 1; j < n - k + 1; j++) {
                    if (num[j] != '.') {
                        if (num[lastNumIdx] > num[j]) {
                            maxNumIdx = lastNumIdx;
                            break;
                        }
                        lastNumIdx = j;
                    }

                    if (num[maxNumIdx] < num[j]) {
                        maxNumIdx = j;
                    }
                }
                num[maxNumIdx] = '.';
                k--;

                // cout << "k=" << k << "," << num << endl;
                if (maxNumIdx == i) i++;
            } else i++;
        }


        // cout << "ans:" << num << endl;

        bool hasHead = false;
        for (int i = 0; i < n; i++)
            if (num[i] > '.')
                if (num[i] != '0' || hasHead) {
                    ss << num[i];
                    hasHead = true;
                }

        if (ss.str().length() == 0) return "0";
        return ss.str();
    }
};
// @lc code=end
