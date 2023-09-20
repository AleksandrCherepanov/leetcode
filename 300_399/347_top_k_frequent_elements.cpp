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

        priority_queue<pair<int, int>> heap;
        for (auto [key, val]: umap) {
            heap.push({-val, key});
            if (heap.size() > k) {
                heap.pop();
            }
        }

        vector<int> result;
        while (heap.size() > 0) {
            result.push_back(heap.top().second);
            heap.pop();
        }

        return result;
    }
};