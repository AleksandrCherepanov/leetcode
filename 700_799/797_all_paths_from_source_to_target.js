var allPathsSourceTarget = function(graph) {
    const ans = [];

    const find = (curr, node) => {
        if (node === graph.length - 1) {
            ans.push([...curr])
            return;
        }

        for (let n of graph[node]) {
            curr.push(n);
            find(curr, n)
            curr.pop();
        }
    }

    find([0], 0);

    return ans;
};
