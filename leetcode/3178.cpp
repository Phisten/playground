class Solution {
public:
    int numberOfChild(int n, int k) {
        int cur = 0;
        int direct = 1;
        while (k > 0) {
          if (cur == n - 1) direct = -1;
          if (cur == 0) direct = 1;
          cur+=direct;
          k--;
        }

        return cur;
    }
};
