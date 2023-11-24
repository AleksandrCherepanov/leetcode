#include <iostream>
#include <vector>
#include <queue>

using namespace std;

int findCheapestPrice(int n, vector<vector<int> >& flights, int src, int dst, int k) {
    vector<vector<pair<int, int>>> graph (n);
    for (vector<int> edge: flights) {
        graph[edge[0]].push_back({edge[1], edge[2]});
    }

    vector<int> steps (n, {INT_MAX});
    priority_queue<vector<int>, vector<vector<int>>, greater<vector<int>>> heap;
    heap.push({0, 0, src});

    while(!heap.empty()) {
        int currPrice = heap.top()[0];
        int currSteps = heap.top()[1];
        int node = heap.top()[2];
        heap.pop();

        if (currSteps > steps[node] || currSteps > k + 1) {
            continue;
        }

        steps[node] = currSteps;

        if (node == dst) {
            return currPrice;
        }

        for(pair<int, int> edge: graph[node]) {
            heap.push({currPrice + edge.second, currSteps + 1, edge.first});                    
        }
    }

    return -1;
}