var maximumDetonation = function(bombs) {
    let graph = new Map();

    const distance = (x1, y1, x2, y2) => {
        return Math.sqrt(Math.pow(x2 - x1, 2) + Math.pow(y2 - y1, 2));
    }

    const label = (a, b) => {
        return [a, b].join(',');
    }

    let result = 0;
    for (let i = 0; i < bombs.length; i++) {
        let stack = [bombs[i]];
        let seen = new Map([[i, true]]);
        let count = 1;

        while (stack.length > 0) {
            const [x1, y1, r1] = stack.pop();
            for (let j = 0; j < bombs.length; j++) {
                const [x2, y2] = bombs[j];
                if (!seen.has(j) && distance(x1,y1,x2,y2) <= r1) {
                    seen.set(j, true);
                    stack.push(bombs[j]);
                    count++;
                }
            }
        }

        result = Math.max(result, count);
    }

    return result;
};
