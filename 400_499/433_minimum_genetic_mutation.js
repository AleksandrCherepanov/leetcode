var minMutation = function(startGene, endGene, bank) {
    let bankMap = new Map();
    for (const b of bank) {
        bankMap.set(b, true);
    }

    if (!bankMap.has(endGene)) {
        return -1;
    }

    let queue = [startGene];
    let seen = new Map();
    let steps = 0;
    seen.set(startGene, true);

    while (queue.length > 0) {
        const currLen = queue.length;
        for (let i = 0; i < currLen; i++) {
            const q = queue.shift();
            if (q === endGene) {
                return steps;
            }
            for (let j = 0; j < 8; j++) {
                for (const c of ['A', 'C', 'G', 'T']) {
                    const n = q.slice(0, j) + c + q.slice(j + 1, 8);
                    if (bankMap.has(n) && !seen.has(n)) {
                        seen.set(n, true)
                        queue.push(n)
                    }
                }
            }
        }
        steps++;
    }

    return -1;
};
