using namespace std;

#include <vector>
#include <queue>

class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        priority_queue<int> minq;
        for (auto n: nums) {
            minq.push(-n);
            if (minq.size() > k) {
                minq.pop();
            }
        }

        return -minq.top();      
    }
};