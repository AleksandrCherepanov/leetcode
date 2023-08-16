using namespace std;
#include <vector>;
#include <queue>;
#include <unordered_map>;

class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int, int> umap;
        for (auto n: nums) {
            umap[n]++;
        }

        priority_queue<int, vector<int>, greater<int>> minq;
        for (auto x: umap) {
            minq.push(x.second);
            if (minq.size() > k) {
                minq.pop();
            }
        }

        vector<int> result;
        while (minq.size() > 0) {
            result.push_back(umap[minq.top()]);
            minq.pop();
        }

        return result;
    }
};