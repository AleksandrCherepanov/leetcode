class Solution {
public:
    int halveArray(vector<int>& nums) {
        priority_queue<float> heap;
        double sum = 0.0f;
        for (int i = 0; i < nums.size(); i++) {
            heap.push((double)nums[i]);
            sum += nums[i];
        }

        auto halfSum = sum / 2;
        int result = 0;

        while (true) {
            result++;
            auto max = heap.top();
            heap.pop();

            max /= 2;
            cout << max << endl;
            sum -= max;

            if (sum <= halfSum) {
                break;
            }

            heap.push(max);
        }

        return result;
    }
};