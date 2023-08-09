/**
 * @param {string[][]} equations
 * @param {number[]} values
 * @param {string[][]} queries
 * @return {number[]}
 */
var calcEquation = function(equations, values, queries) {
    let graph = new Map();
    for (let i= 0; i < equations.length; i++) {
        if (graph.has(equations[i][0])) {
            graph.get(equations[i][0]).push({node: equations[i][1], cost: values[i], op: '*'});
        } else {
            graph.set(equations[i][0], [{node: equations[i][1], cost: values[i], op: '*'}]);
        }
        if (graph.has(equations[i][1])) {
            graph.get(equations[i][1]).push({node: equations[i][0], cost: values[i], op: '/'});
        } else {
            graph.set(equations[i][1], [{node: equations[i][0], cost: values[i], op: '/'}]);
        }
    }

    let ans = new Array(queries.length).fill(-1);
    for (let j = 0; j < queries.length; j++) {
        const q = queries[j];
        let seen = new Map();
        let queue = [{node: q[0], cost: 1}];
        seen.set(q[0], true);

        while (queue.length > 0) {
            const currLen = queue.length;
            for (let i = 0; i < currLen; i++) {
                const v = queue.shift();
                if (graph.has(v.node)) {
                    if (v.node === q[1]) {
                        ans[j] = v.cost;
                        break;
                    }


                    for (const n of graph.get(v.node)) {
                        if (!seen.has(n.node)) {
                            seen.set(n.node, true)
                            queue.push({node: n.node, cost: n.op === '*' ? v.cost * n.cost : v.cost / n.cost});
                        }
                    }
                }
            }
        }
    }

    return ans;
};
