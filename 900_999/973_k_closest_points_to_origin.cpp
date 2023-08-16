using namespace std;

#include <vector>;
#include <queue>;

class Solution {
public:
    vector<vector<int>> kClosest(vector<vector<int>>& points, int k) {
        priority_queue <pair<double, pair<int, int>>> maxq;

        vector<vector<int>> ans;

        for (auto p: points) {
            int dx = p[0] - 0;
            int dy = p[1] - 0;
            double d = sqrt(dx * dx + dy * dy);
            
            maxq.push({d, {p[0], p[1]}});
            if (maxq.size() > k) {
                maxq.pop();
            }
        }

        while (maxq.size() > 0) {
            ans.push_back({maxq.top().second.first, maxq.top().second.second});
            maxq.pop();
        }


        return ans;
    }
};