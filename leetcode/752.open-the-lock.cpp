/*
 * @lc app=leetcode id=752 lang=cpp
 *
 * [752] Open the Lock
 */

// @lc code=start
class Solution {
public:
    string addOne(string src, int index) {
        if (src[index] == '9') src[index] = '0';
        else src[index] += 1;
        return src;
    }
    string subOne(string src, int index) {
        if (src[index] == '0') src[index] = '9';
        else src[index] -= 1;
        return src; 
    }
    int openLock(vector<string>& deadends, string target) {
        queue<string> q;

        unordered_map<string, bool> ht;
        for (int i = 0; i < deadends.size(); i++) {
            ht[deadends[i]] = true;
        }

        //check deadends
        int deadCount = 0;
        for (int i = 0; i < 4; i++) {
            string addStr = addOne(target, i);
            if (ht[addStr] == true) deadCount++;
            string subStr = subOne(target, i);
            if (ht[subStr] == true) deadCount++;
        }
        if (deadCount == 8) return -1;
        if (ht["0000"] == true) return -1;


        int ans = 0;
        string lastStepStr = "0000";
        q.push("0000");
        ht["0000"] = true;

        int nextLevelCount = 0;
        int levelCount = 1;
        while (q.size() > 0) {
            string cur = q.front();
            q.pop();
            levelCount--;

            // cout << cur << " step:" << ans << endl;

            if (cur == target) return ans;

            for (int i = 0; i < 4; i++) {
                string addStr = addOne(cur, i);
                if (ht[addStr] != true) {
                    q.push(addStr);
                    ht[addStr] = true;
                    nextLevelCount++;
                }
                
                string subStr = subOne(cur, i);
                if (ht[subStr] != true) {
                    q.push(subStr);
                    ht[subStr] = true;
                    nextLevelCount++;
                }
            }

            if (levelCount == 0) {
                ans++;
                levelCount = nextLevelCount;
                nextLevelCount = 0;
            }
        }

        return -1;
    }
};
// @lc code=end

