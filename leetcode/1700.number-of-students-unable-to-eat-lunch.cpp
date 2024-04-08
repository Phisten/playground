/*
 * @lc app=leetcode id=1700 lang=cpp
 *
 * [1700] Number of Students Unable to Eat Lunch
 */

// @lc code=start
class Solution {
public:
    int countStudents(vector<int>& students, vector<int>& sandwiches) {
        int n = students.size();

        int need0 = 0;
        int need1 = 0;

        for (int i = 0; i < n; i++)
            if (students[i] == 0) need0++;
            else need1++;
        for (int i = 0; i < n; i++)
            if (sandwiches[i] == 0)
                if (need0 > 0) need0--;
                else return need1;
            else
                if (need1 > 0) need1--;
                else return need0;

        return 0;
    }
};
// @lc code=end

