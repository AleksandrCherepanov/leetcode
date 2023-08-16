using namespace std;

#include <vector>;
#include <queue>;

class Solution {
public:
    int connectSticks(vector<int>& sticks) {
        priority_queue<int, vector<int>, greater<int>> minq;
        for (auto s: sticks) {
            minq.push(s);
        }

        int cost = 0;
        while (minq.size() > 1) {
            auto s1 = minq.top();
            minq.pop();

            auto s2 = minq.top();
            minq.pop();
        }

        return cost;
    }
};