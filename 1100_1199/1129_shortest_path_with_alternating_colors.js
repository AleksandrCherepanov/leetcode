/**
 * @param {number} n
 * @param {number[][]} redEdges
 * @param {number[][]} blueEdges
 * @return {number[]}
 */
var shortestAlternatingPaths = function(n, redEdges, blueEdges) {
    let addToGraph = (color, edges) => {
        for (let i = 0; i < n; i++) {
            graph.get(color).set(i, []);
        }

        for (const [x, y] of edges) {
            graph.get(color).get(x).push(y);
        }
    }

    const RED = 0;
    const BLUE = 1;

    let graph = new Map();
    graph.set(RED, new Map());
    graph.set(BLUE, new Map());
    addToGraph(RED, redEdges);
    addToGraph(BLUE, blueEdges);

    let ans = [];
    for (let i = 0; i < n; i++) {
        ans.push(Infinity);
    }

    let seen = [];
    for (let i = 0; i < n; i++) {
        seen.push(new Array(2).fill(false));
    }

    let queue = [[0, RED], [0, BLUE]];
    seen[0][RED] = true;
    seen[0][BLUE] = true;

    let steps = 0;

    while (queue.length > 0) {
        currLen = queue.length;
        for (let i = 0; i < currLen; i++) {
            let [node, color] = queue.shift();
            ans[node] = Math.min(ans[node], steps);

            for (const n of graph.get(color).get(node)) {
                if (!seen[n][1-color]) {
                    seen[n][1-color] = true;
                    queue.push([n, 1 -color]);
                }
            }
        }

        steps++;
    }

    return ans.map((el) => { return el === Infinity ? -1 : el; });
};
