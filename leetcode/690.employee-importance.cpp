/*
 * @lc app=leetcode id=690 lang=cpp
 *
 * [690] Employee Importance
 */

// @lc code=start
/*
// Definition for Employee.
class Employee {
public:
    int id;
    int importance;
    vector<int> subordinates;
};
*/
/*
// Definition for Employee.
class Employee {
public:
    int id;
    int importance;
    vector<int> subordinates;
};
*/

class Solution {
public:
    int getImportance(vector<Employee*> employees, int id) {
        stack<Employee*> st;

        unordered_map<int, Employee*> ht;

        int n = employees.size();
        for (int i = 0; i < n; i++) {
            ht[employees[i]->id] = employees[i];
        }

        st.push(ht[id]);
        int ans = 0;
        while (st.size() > 0) {
            Employee* cur = st.top();
            st.pop();

            ans += cur->importance;

            int nn = cur->subordinates.size();
            for (int i = 0; i < nn; i++) {
                st.push(ht[cur->subordinates[i]]);
            }
        }

        return ans;
    }
};
// @lc code=end

