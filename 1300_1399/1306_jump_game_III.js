var canReach = function(arr, start) {
    let seen = new Map();
    seen.set(start.toString(), true);

    let stack = [start];
    while (stack.length > 0) {
        let currIdx = stack.pop();
        if (arr[currIdx] === 0) {
            return true;
        }
        for (const newIdx of [currIdx + arr[currIdx], currIdx - arr[currIdx]]) {
            if (newIdx >= 0 && newIdx < arr.length && !seen.has(newIdx.toString())) {
                stack.push(newIdx);
                seen.set(newIdx.toString(), true);
            }
        }
    }

    return false;
};
