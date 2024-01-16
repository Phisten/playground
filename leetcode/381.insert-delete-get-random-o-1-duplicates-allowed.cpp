/*
 * @lc app=leetcode id=381 lang=cpp
 *
 * [381] Insert Delete GetRandom O(1) - Duplicates allowed
 */

// @lc code=start
#include <unordered_map>
#include <vector>
#include <iostream>
using namespace std;
class RandomizedCollection
{
    unordered_map<int, vector<int>> ht;
    vector<int> vt;

public:
    RandomizedCollection()
    {
    }

    bool insert(int val)
    {
        bool notfound = ht.count(val) == 0;

        ht[val].push_back(vt.size());
        vt.push_back(val);

        cout << "add, size=" << vt.size() << endl;

        return notfound;
    }

    bool remove(int val)
    {
        auto find = ht.count(val) > 0;
        if (find)
        {
            auto vtLastVal = vt.back();
            auto removeIdx = ht[val].back();

            ht[val].pop_back();
            if (ht[val].size() == 0)
                ht.erase(val);

            if (vtLastVal != val)
            {
                auto vtLastInHt = ht[vtLastVal];
                vtLastInHt[vtLastInHt.size() - 1] = removeIdx;
            }

            cout << "remove, val=" << val << ", moveLast:" << vtLastVal << " to:" << removeIdx << endl;

            vt[removeIdx] = vtLastVal;
            vt.pop_back();
        }
        for (const auto &value : vt)
        {
            cout << value << ",";
        }
        cout << endl;

        return find;
    }

    int getRandom()
    {
        return vt[rand() % vt.size()];
    }
};

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * RandomizedCollection* obj = new RandomizedCollection();
 * bool param_1 = obj->insert(val);
 * bool param_2 = obj->remove(val);
 * int param_3 = obj->getRandom();
 */
// @lc code=end
