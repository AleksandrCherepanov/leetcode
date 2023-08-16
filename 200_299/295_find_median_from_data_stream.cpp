class MedianFinder {
public:
    priority_queue<int, vector<int>, greater<int>> minq;
    priority_queue<int> maxq;
    MedianFinder() {

    }

    void addNum(int num) {
        maxq.push(num);
        minq.push(maxq.top());
        maxq.pop();

        if (minq.size() > maxq.size()) {
            maxq.push(minq.top());
            minq.pop();
        }
    }

    double findMedian() {
        if (maxq.size() > minq.size()) {
            return maxq.top();
        }

        return (minq.top() + maxq.top()) / 2.0;
    }
};
