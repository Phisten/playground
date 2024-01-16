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
    struct VtVal
    {
        int val;
        int idxInHt;
    };
    typedef vector<int> IdxInVt;

    // key=Val, values=index in vt
    unordered_map<int, IdxInVt>
        ht;
    // value pair = (Val, index in ht[Val])
    vector<VtVal> vt;

public:
    RandomizedCollection()
    {
    }

    bool insert(int val)
    {
        bool notfound = ht.count(val) == 0;

        ht[val].push_back(vt.size());
        vt.push_back(VtVal{val : val, idxInHt : (int)(ht[val].size()) - 1});

        return notfound;
    }

    bool remove(int val)
    {
        auto find = ht.count(val) > 0;
        if (find)
        {
            auto vtLast = vt.back();
            auto vtRemoveIdx = ht[val].back();

            ht[val].pop_back();
            if (ht[val].size() == 0)
                ht.erase(val);

            if (vtLast.val != val)
            {
                ht[vtLast.val][vtLast.idxInHt] = vtRemoveIdx;
            }

            vt[vtRemoveIdx].val = vtLast.val;
            vt[vtRemoveIdx].idxInHt = vtLast.idxInHt;
            vt.pop_back();
        }

        return find;
    }

    int getRandom()
    {
        return vt[rand() % vt.size()].val;
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
