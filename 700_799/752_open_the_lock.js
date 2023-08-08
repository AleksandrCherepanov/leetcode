/**
 * @param {string[]} deadends
 * @param {string} target
 * @return {number}
 */
var openLock = function(deadends, target) {
    if (target === '0000') {
        return 0;
    }

    let seen = new Map();
    let deadMap = new Map();
    for (const d of deadends) {
        deadMap.set(d, true);
    }

    if (deadMap.has('0000')) {
        return -1;
    }

    let queue = ['0000'];
    let steps = 0;
    seen.set('0000', true);

    while (queue.length > 0) {
        const currLen = queue.length;

        for (let k = 0; k < currLen; k++) {
            let state = queue.shift();
            if (state === target) {
                return steps;
            }

            for (let i = 0; i < 4; i++) {
                for (const j of [1, -1]) {
                    let wheel = state[i];
                    const numWheel = (+wheel + j + 10) % 10;
                    const newState = state.substring(0, i) + numWheel.toString() + state.substring(i + 1, 4);
                    if (!seen.has(newState) && !deadMap.has(newState)) {
                        seen.set(newState, true);
                        queue.push(newState);
                    }
                }
            }
        }

        steps++;
    }

    return -1;
};
