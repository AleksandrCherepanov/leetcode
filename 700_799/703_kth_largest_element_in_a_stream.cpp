using namespace std;

#include <vector>
#include <queue>

class KthLargest {
public:
    int k;
    priority_queue<int, vector<int>, greater<int>> minq;
    KthLargest(int k, vector<int>& nums) {
        this->k = k;
        for (auto n: nums) {
            minq.push(n);
        }

        while (minq.size() > this->k) {
            minq.pop();
        }
    }
    
    int add(int val) {
        minq.push(val);
        if (minq.size() > this->k) {
            minq.pop();
        }

        return minq.top();
    }
};